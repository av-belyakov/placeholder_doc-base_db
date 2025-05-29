// Пакет confighandler формирует конфигурационные настройки приложения
package confighandler

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

func New(rootDir string) (*ConfigApp, error) {
	conf := &ConfigApp{}

	var (
		validate *validator.Validate
		envList  map[string]string = map[string]string{
			"GO_PHDOCBASEDB_MAIN": "",

			//Подключение к NATS
			"GO_PHDOCBASEDB_NHOST":        "",
			"GO_PHDOCBASEDB_NPORT":        "",
			"GO_PHDOCBASEDB_NCOMMAND":     "",
			"GO_PHDOCBASEDB_REQUESTS":     "",
			"GO_PHDOCBASEDB_NCACHETTL":    "",
			"GO_PHDOCBASEDB_NSUBLISTENER": "",

			//Настройки доступа к БД в которую будут записыватся alert и case
			"GO_PHDOCBASEDB_DBSTORAGEN":      "",
			"GO_PHDOCBASEDB_DBSTORAGEHOST":   "",
			"GO_PHDOCBASEDB_DBSTORAGEPORT":   "",
			"GO_PHDOCBASEDB_DBSTORAGENAME":   "",
			"GO_PHDOCBASEDB_DBSTORAGEUSER":   "",
			"GO_PHDOCBASEDB_DBSTORAGEPASSWD": "",

			//Настройки доступа к БД в которую будут записыватся логи
			"GO_PHDOCBASEDB_DBWLOGHOST":        "",
			"GO_PHDOCBASEDB_DBWLOGPORT":        "",
			"GO_PHDOCBASEDB_DBWLOGNAME":        "",
			"GO_PHDOCBASEDB_DBWLOGUSER":        "",
			"GO_PHDOCBASEDB_DBWLOGPASSWD":      "",
			"GO_PHDOCBASEDB_DBWLOGSTORAGENAME": "",
		}
	)

	getFileName := func(sf, confPath string, lfs []fs.DirEntry) (string, error) {
		for _, v := range lfs {
			if v.Name() == sf && !v.IsDir() {
				return filepath.Join(confPath, v.Name()), nil
			}
		}

		return "", fmt.Errorf("file '%s' is not found", sf)
	}

	setCommonSettings := func(fn string) error {
		viper.SetConfigFile(fn)
		viper.SetConfigType("yml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		ls := Logs{}
		if ok := viper.IsSet("LOGGING"); ok {
			if err := viper.GetViper().Unmarshal(&ls); err != nil {
				return err
			}

			conf.Common.Logs = ls.Logging
		}

		z := ZabbixSet{}
		if ok := viper.IsSet("ZABBIX"); ok {
			if err := viper.GetViper().Unmarshal(&z); err != nil {
				return err
			}

			np := 10051
			if z.Zabbix.NetworkPort != 0 && z.Zabbix.NetworkPort < 65536 {
				np = z.Zabbix.NetworkPort
			}

			conf.Common.Zabbix = ZabbixOptions{
				NetworkPort: np,
				NetworkHost: z.Zabbix.NetworkHost,
				ZabbixHost:  z.Zabbix.ZabbixHost,
				EventTypes:  z.Zabbix.EventTypes,
			}
		}

		return nil
	}

	setSpecial := func(fn string) error {
		viper.SetConfigFile(fn)
		viper.SetConfigType("yml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		//Настройки для модуля подключения к NATS
		if viper.IsSet("NATS.host") {
			conf.NATS.Host = viper.GetString("NATS.host")
		}
		if viper.IsSet("NATS.port") {
			conf.NATS.Port = viper.GetInt("NATS.port")
		}
		if viper.IsSet("NATS.cache_ttl") {
			conf.NATS.CacheTTL = viper.GetInt("NATS.cache_ttl")
		}
		if viper.IsSet("NATS.command") {
			conf.NATS.Command = viper.GetString("NATS.command")
		}
		if viper.IsSet("NATS.requests") {
			conf.NATS.Requests = viper.GetStringMapString("NATS.requests")
		}
		if viper.IsSet("NATS.subscriptions") {
			conf.NATS.Subscriptions = viper.GetStringMapString("NATS.subscriptions")
		}

		// Настройки доступа к БД в которую будет записыватся основная информация
		if viper.IsSet("DATABASESTORAGE.host") {
			conf.StorageDB.Host = viper.GetString("DATABASESTORAGE.host")
		}
		if viper.IsSet("DATABASESTORAGE.port") {
			conf.StorageDB.Port = viper.GetInt("DATABASESTORAGE.port")
		}
		if viper.IsSet("DATABASESTORAGE.user") {
			conf.StorageDB.User = viper.GetString("DATABASESTORAGE.user")
		}
		if viper.IsSet("DATABASESTORAGE.namedb") {
			conf.StorageDB.NameDB = viper.GetString("DATABASESTORAGE.namedb")
		}
		if viper.IsSet("DATABASESTORAGE.storage_name_db") {
			conf.StorageDB.Storage = viper.GetStringMapString("DATABASESTORAGE.storage_name_db")
		}

		// Настройки доступа к БД в которую будут записыватся логи
		if viper.IsSet("DATABASEWRITELOG.host") {
			conf.LogDB.Host = viper.GetString("DATABASEWRITELOG.host")
		}
		if viper.IsSet("DATABASEWRITELOG.port") {
			conf.LogDB.Port = viper.GetInt("DATABASEWRITELOG.port")
		}
		if viper.IsSet("DATABASEWRITELOG.user") {
			conf.LogDB.User = viper.GetString("DATABASEWRITELOG.user")
		}
		if viper.IsSet("DATABASEWRITELOG.namedb") {
			conf.LogDB.NameDB = viper.GetString("DATABASEWRITELOG.namedb")
		}
		if viper.IsSet("DATABASEWRITELOG.storage_name_db") {
			conf.LogDB.StorageNameDB = viper.GetString("DATABASEWRITELOG.storage_name_db")
		}

		return nil
	}

	validate = validator.New(validator.WithRequiredStructEnabled())

	for v := range envList {
		if env, ok := os.LookupEnv(v); ok {
			envList[v] = env
		}
	}

	rootPath, err := supportingfunctions.GetRootPath(rootDir)
	if err != nil {
		return conf, err
	}

	confPath := filepath.Join(rootPath, "config")
	list, err := os.ReadDir(confPath)
	if err != nil {
		return conf, err
	}

	fileNameCommon, err := getFileName("config.yml", confPath, list)
	if err != nil {
		return conf, err
	}

	//читаем общий конфигурационный файл
	if err := setCommonSettings(fileNameCommon); err != nil {
		return conf, err
	}

	var fn string
	if envList["GO_PHDOCBASEDB_MAIN"] == "development" {
		fn, err = getFileName("config_dev.yml", confPath, list)
		if err != nil {
			return conf, err
		}
	} else if envList["GO_PHDOCBASEDB_MAIN"] == "test" {
		fn, err = getFileName("config_test.yml", confPath, list)
		if err != nil {
			return conf, err
		}
	} else {
		fn, err = getFileName("config_prod.yml", confPath, list)
		if err != nil {
			return conf, err
		}
	}

	if err := setSpecial(fn); err != nil {
		return conf, err
	}

	//Настройки для модуля подключения к NATS
	if envList["GO_PHDOCBASEDB_NHOST"] != "" {
		conf.NATS.Host = envList["GO_PHDOCBASEDB_NHOST"]
	}
	if envList["GO_PHDOCBASEDB_NPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHDOCBASEDB_NPORT"]); err == nil {
			conf.NATS.Port = p
		}
	}
	if envList["GO_PHDOCBASEDB_NCACHETTL"] != "" {
		if ttl, err := strconv.Atoi(envList["GO_PHDOCBASEDB_NCACHETTL"]); err == nil {
			conf.NATS.CacheTTL = ttl
		}
	}

	if envList["GO_PHDOCBASEDB_NCOMMAND"] != "" {
		conf.NATS.Command = envList["GO_PHDOCBASEDB_NCOMMAND"]
	}
	if envList["GO_PHDOCBASEDB_REQUESTS"] != "" {
		reqlistener := envList["GO_PHDOCBASEDB_REQUESTS"]
		if !strings.Contains(reqlistener, ";") {
			if tmp := strings.Split(reqlistener, ":"); len(tmp) == 2 {
				conf.NATS.Requests[tmp[0]] = tmp[1]
			}
		} else {
			for sl := range strings.SplitSeq(reqlistener, ";") {
				if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
					conf.NATS.Requests[tmp[0]] = tmp[1]
				}
			}
		}
	}
	if envList["GO_PHDOCBASEDB_NSUBLISTENER"] != "" {
		sublistener := envList["GO_PHDOCBASEDB_NSUBLISTENER"]
		if !strings.Contains(sublistener, ";") {
			if tmp := strings.Split(sublistener, ":"); len(tmp) == 2 {
				conf.NATS.Subscriptions[tmp[0]] = tmp[1]
			}
		} else {
			for sl := range strings.SplitSeq(sublistener, ";") {
				if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
					conf.NATS.Subscriptions[tmp[0]] = tmp[1]
				}
			}
		}
	}

	//Настройки доступа к БД в которую будет добавлятся информация по alert и case
	if envList["GO_PHDOCBASEDB_DBSTORAGEHOST"] != "" {
		conf.StorageDB.Host = envList["GO_PHDOCBASEDB_DBSTORAGEHOST"]
	}
	if envList["GO_PHDOCBASEDB_DBSTORAGEPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHDOCBASEDB_DBSTORAGEPORT"]); err == nil {
			conf.StorageDB.Port = p
		}
	}
	if envList["GO_PHDOCBASEDB_DBSTORAGENAME"] != "" {
		conf.StorageDB.NameDB = envList["GO_PHDOCBASEDB_DBSTORAGENAME"]
	}
	if envList["GO_PHDOCBASEDB_DBSTORAGEUSER"] != "" {
		conf.StorageDB.User = envList["GO_PHDOCBASEDB_DBSTORAGEUSER"]
	}
	if envList["GO_PHDOCBASEDB_DBSTORAGEPASSWD"] != "" {
		conf.StorageDB.Passwd = envList["GO_PHDOCBASEDB_DBSTORAGEPASSWD"]
	}
	if envList["GO_PHDOCBASEDB_DBSTORAGEN"] != "" {
		sublistener := envList["GO_PHDOCBASEDB_DBSTORAGEN"]
		if !strings.Contains(sublistener, ";") {
			if tmp := strings.Split(sublistener, ":"); len(tmp) == 2 {
				conf.StorageDB.Storage[tmp[0]] = tmp[1]
			}
		} else {
			for _, sl := range strings.Split(sublistener, ";") {
				if tmp := strings.Split(sl, ":"); len(tmp) == 2 {
					conf.StorageDB.Storage[tmp[0]] = tmp[1]
				}
			}
		}
	}

	//Настройки доступа к БД в которую будут записыватся логи
	if envList["GO_PHDOCBASEDB_DBWLOGHOST"] != "" {
		conf.LogDB.Host = envList["GO_PHDOCBASEDB_DBWLOGHOST"]
	}
	if envList["GO_PHDOCBASEDB_DBWLOGPORT"] != "" {
		if p, err := strconv.Atoi(envList["GO_PHDOCBASEDB_DBWLOGPORT"]); err == nil {
			conf.LogDB.Port = p
		}
	}
	if envList["GO_PHDOCBASEDB_DBWLOGNAME"] != "" {
		conf.LogDB.NameDB = envList["GO_PHDOCBASEDB_DBWLOGNAME"]
	}
	if envList["GO_PHDOCBASEDB_DBWLOGUSER"] != "" {
		conf.LogDB.User = envList["GO_PHDOCBASEDB_DBWLOGUSER"]
	}
	if envList["GO_PHDOCBASEDB_DBWLOGPASSWD"] != "" {
		conf.LogDB.Passwd = envList["GO_PHDOCBASEDB_DBWLOGPASSWD"]
	}
	if envList["GO_PHDOCBASEDB_DBWLOGSTORAGENAME"] != "" {
		conf.LogDB.StorageNameDB = envList["GO_PHDOCBASEDB_DBWLOGSTORAGENAME"]
	}

	//выполняем проверку заполненой структуры
	if err = validate.Struct(conf); err != nil {
		return conf, err
	}

	return conf, nil
}
