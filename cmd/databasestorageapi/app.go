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
	dbs := &DatabaseStorage{counter: counter, logger: logger}

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
