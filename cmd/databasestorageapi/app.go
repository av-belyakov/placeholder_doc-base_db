package databasestorageapi

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// New настраивает новый модуль взаимодействия с API Database
func New(counter interfaces.Counter, logger interfaces.Logger, opts ...DatabaseStorageOptions) (*DatabaseStorage, error) {
	dbs := &DatabaseStorage{
		counter:  counter,
		logger:   logger,
		chInput:  make(chan SettingsChanInput),
		chOutput: make(chan SettingsChanOutput),
	}

	for _, opt := range opts {
		if err := opt(dbs); err != nil {
			return dbs, err
		}
	}

	return dbs, nil
}

// Start инициализирует новый модуль взаимодействия с API Database
func (dbs *DatabaseStorage) Start(ctx context.Context) error {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:%d", dbs.settings.host, dbs.settings.port)},
		Username:  dbs.settings.user,
		Password:  dbs.settings.passwd,
		Transport: &http.Transport{
			MaxIdleConns:          10,              //число открытых TCP-соединений, которые в данный момент не используются
			IdleConnTimeout:       1 * time.Second, //время, через которое закрываются такие неактивные соединения
			MaxIdleConnsPerHost:   10,              //число неактивных TCP-соединений, которые допускается устанавливать на один хост
			ResponseHeaderTimeout: 2 * time.Second, //время в течении которого сервер ожидает получение ответа после записи заголовка запроса
			DialContext: (&net.Dialer{
				Timeout: 3 * time.Second,
				//KeepAlive: 1 * time.Second,
			}).DialContext,
		},
		//RetryOnError: ,
		//RetryOnStatus: ,
	})
	if err != nil {
		return supportingfunctions.CustomError(err)
	}

	dbs.client = es

	go dbs.router(ctx)

	return nil
}

/*
fatal error: concurrent map iteration and map write

goroutine 122881 [running]:
internal/runtime/maps.fatal({0xe1f153?, 0x474d6f?})
	/usr/local/go/src/runtime/panic.go:1058 +0x18
internal/runtime/maps.(*Iter).Next(0xc0002f3308?)
	/usr/local/go/src/internal/runtime/maps/table.go:683 +0x86
github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi.(*DatabaseStorage).InsertDocument(0xc000398000, {0xf27108, 0xc0000adbc0}, {0xc0000d9200, 0x28}, {0xc0002f4000, 0x136c, 0x1500})
	/go/src/c27a435c/cmd/databasestorageapi/requestshandlers.go:158 +0x365
github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi.(*DatabaseStorage).addAlert(0xc000398000, {0xf27108, 0xc0000adbc0}, {0xd73f20, 0xc000004808})
	/go/src/c27a435c/cmd/databasestorageapi/addalerthandler.go:120 +0xcca
created by github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi.(*DatabaseStorage).router in goroutine 38
	/go/src/c27a435c/cmd/databasestorageapi/routers.go:44 +0x60a
*/
