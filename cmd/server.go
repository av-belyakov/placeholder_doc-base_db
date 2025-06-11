package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/av-belyakov/simplelogger"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/elasticsearchapi"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/natsapi"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/wrappers"
	"github.com/av-belyakov/placeholder_doc-base_db/constants"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/confighandler"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/countermessage"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/logginghandler"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

func server(ctx context.Context) {
	var nameRegionalObject string
	if os.Getenv("GO_PHDOCBASEDB_MAIN") == "development" {
		nameRegionalObject = "gcm-dev"
	} else if os.Getenv("GO_PHDOCBASEDB_MAIN") == "test" {
		nameRegionalObject = "gcm-test"
	} else {
		nameRegionalObject = "gcm"
	}

	rootPath, err := supportingfunctions.GetRootPath(constants.Root_Dir)
	if err != nil {
		log.Fatalf("error, it is impossible to form root path (%s)", err.Error())
	}

	// ****************************************************************************
	// *********** инициализируем модуль чтения конфигурационного файла ***********
	conf, err := confighandler.New(rootPath)
	if err != nil {
		log.Fatalf("error module 'confighandler': %v", err)
	}

	// ****************************************************************************
	// ********************* инициализация модуля логирования *********************
	var listLog []simplelogger.OptionsManager
	for _, v := range conf.GetListLogs() {
		listLog = append(listLog, v)
	}
	opts := simplelogger.CreateOptions(listLog...)
	simpleLogger, err := simplelogger.NewSimpleLogger(ctx, constants.Root_Dir, opts)
	if err != nil {
		log.Fatalf("error module 'simplelogger': %v", err)
	}

	//*********************************************************************************
	//********** инициализация модуля взаимодействия с БД для передачи логов **********
	confDB := conf.GetLogDB()
	if esc, err := elasticsearchapi.NewElasticsearchConnect(elasticsearchapi.Settings{
		Port:               confDB.Port,
		Host:               confDB.Host,
		User:               confDB.User,
		Passwd:             confDB.Passwd,
		IndexDB:            confDB.StorageNameDB,
		NameRegionalObject: nameRegionalObject,
	}); err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())
	} else {
		//подключение логирования в БД
		simpleLogger.SetDataBaseInteraction(esc)
	}

	// ************************************************************************
	// ************* инициализация модуля взаимодействия с Zabbix *************
	chZabbix := make(chan interfaces.Messager)
	confZabbix := conf.GetZabbix()
	wziSettings := wrappers.WrappersZabbixInteractionSettings{
		NetworkPort: confZabbix.NetworkPort,
		NetworkHost: confZabbix.NetworkHost,
		ZabbixHost:  confZabbix.ZabbixHost,
	}
	eventTypes := []wrappers.EventType(nil)
	for _, v := range confZabbix.EventTypes {
		eventTypes = append(eventTypes, wrappers.EventType{
			IsTransmit: v.IsTransmit,
			EventType:  v.EventType,
			ZabbixKey:  v.ZabbixKey,
			Handshake: wrappers.Handshake{
				TimeInterval: v.Handshake.TimeInterval,
				Message:      v.Handshake.Message,
			},
		})
	}
	wziSettings.EventTypes = eventTypes
	wrappers.WrappersZabbixInteraction(ctx, wziSettings, simpleLogger, chZabbix)

	//***************************************************************************
	//************** инициализация обработчика логирования данных ***************
	//фактически это мост между simpleLogger и пакетом соединения с Zabbix
	logging := logginghandler.New(simpleLogger, chZabbix)
	logging.Start(ctx)

	// ***************************************************************************
	// *********** инициализируем модуль счётчика для подсчёта сообщений *********
	counting := countermessage.New(chZabbix)
	counting.Start(ctx)

	// ***********************************************************************
	// ************** инициализация модуля взаимодействия с NATS *************
	confNats := conf.NATS
	apiNats, err := natsapi.New(
		counting,
		logging,
		natsapi.WithHost(confNats.Host),
		natsapi.WithPort(confNats.Port),
		natsapi.WithCacheTTL(confNats.CacheTTL),
		natsapi.WithRequests(confNats.Requests),
		natsapi.WithSendCommand(confNats.Command),
		natsapi.WithSubscriptions(confNats.Subscriptions))
	if err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())

		log.Fatal(err)
	}
	//--- старт модуля
	if err = apiNats.Start(ctx); err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())

		log.Fatal(err)
	}

	// *********************************************************************
	// ************** инициализация модуля взаимодействия с БД *************
	confStorageDB := conf.GetStorageDB()
	apiDBS, err := databasestorageapi.New(
		counting,
		logging,
		databasestorageapi.WithHost(confStorageDB.Host),
		databasestorageapi.WithPort(confStorageDB.Port),
		databasestorageapi.WithNameDB(confStorageDB.NameDB),
		databasestorageapi.WithUser(confStorageDB.User),
		databasestorageapi.WithPasswd(confStorageDB.Passwd),
		databasestorageapi.WithStorage(confStorageDB.Storage))
	if err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())

		log.Fatal(err)
	}
	//--- старт модуля
	if err := apiDBS.Start(ctx); err != nil {
		_ = simpleLogger.Write("error", supportingfunctions.CustomError(err).Error())

		log.Fatal(err)
	}

	// *********************************************************
	// ************** инициализация маршрутизатора *************
	r := NewRouter(
		counting,
		logging,
		ApplicationRouterSettings{
			ChanToNats:   apiNats.GetChanDataToModule(),
			ChanFromNats: apiNats.GetChanDataFromModule(),
			ChanToDBS:    apiDBS.GetChanDataToModule(),
			ChanFromDBS:  apiDBS.GetChanDataFromModule(),
		})
	r.Router(ctx)

	//информационное сообщение
	infoMsg := getInformationMessage(conf)
	_ = simpleLogger.Write("info", infoMsg)

	//для отладки через pprof (только для теста)
	//http://conf.Common.Profiling.Host:conf.Common.Profiling.Port/debug/pprof/
	//go tool pprof http://host:port/debug/pprof/heap
	//go tool pprof http://host:port/debug/pprof/allocs
	//go tool pprof http://host:port/debug/pprof/goroutine
	if os.Getenv("GO_PHDOCBASEDB_MAIN") == "test" || os.Getenv("GO_PHDOCBASEDB_MAIN") == "development" {
		if conf.Common.Profiling.Port > 0 {
			go func() {
				log.Println(http.ListenAndServe(fmt.Sprintf("%s:%d", conf.Common.Profiling.Host, conf.Common.Profiling.Port), nil))
			}()
		}
	}

	<-ctx.Done()
}
