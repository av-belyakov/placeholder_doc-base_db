// Модуль для взаимодействия с API NATS
package natsapi

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/av-belyakov/placeholder_doc-base_db/constants"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// New настраивает новый модуль взаимодействия с API NATS
func New(logger interfaces.Logger, opts ...NatsApiOptions) (*apiNatsModule, error) {
	api := &apiNatsModule{
		cachettl: 10,
		//для логирования
		logger: logger,
		//прием запросов в NATS
		chFromModule: make(chan SettingsOutputChan),
		//передача запросов из NATS
		chToModule: make(chan SettingsInputChan),
	}

	for _, opt := range opts {
		if err := opt(api); err != nil {
			return api, err
		}
	}

	return api, nil
}

// Start инициализирует новый модуль взаимодействия с API NATS
// при инициализации возращается канал для взаимодействия с модулем, все
// запросы к модулю выполняются через данный канал
func (api *apiNatsModule) Start(ctx context.Context) (chan<- SettingsInputChan, <-chan SettingsOutputChan, error) {
	if ctx.Err() != nil {
		return api.chToModule, api.chFromModule, ctx.Err()
	}

	nc, err := nats.Connect(
		fmt.Sprintf("%s:%d", api.host, api.port),
		//имя клиента
		nats.Name(fmt.Sprintf("placeholder_docbase_db.%s", api.nameRegionalObject)),
		//неограниченное количество попыток переподключения
		nats.MaxReconnects(-1),
		//время ожидания до следующей попытки переподключения (по умолчанию 2 сек.)
		nats.ReconnectWait(3*time.Second),
		//обработка разрыва соединения с NATS
		nats.DisconnectErrHandler(func(c *nats.Conn, err error) {
			api.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("the connection with NATS has been disconnected (%w)", err)).Error())
		}),
		//обработка переподключения к NATS
		nats.ReconnectHandler(func(c *nats.Conn) {
			api.logger.Send("info", "the connection to NATS has been re-established")
		}),
		//поиск медленных получателей (не обязательный для данного приложения параметр)
		nats.ErrorHandler(func(c *nats.Conn, s *nats.Subscription, err error) {
			if err == nats.ErrSlowConsumer {
				pendingMsgs, _, err := s.Pending()
				if err != nil {
					api.logger.Send("warning", fmt.Sprintf("couldn't get pending messages: %v", err))

					return
				}

				api.logger.Send("warning", fmt.Sprintf("Falling behind with %d pending messages on subject %q.\n", pendingMsgs, s.Subject))
			}
		}))
	if err != nil {
		return api.chToModule, api.chFromModule, supportingfunctions.CustomError(err)
	}

	log.Printf("%vconnect to NATS with address %v%s:%d%v\n", constants.Ansi_Bright_Green, constants.Ansi_Dark_Gray, api.host, api.port, constants.Ansi_Reset)

	api.natsConn = nc

	//обработчик подписок
	go api.subscriptionHandler()

	//обработчик информации полученной изнутри приложения
	go api.incomingInformationHandler(ctx)

	go func(ctx context.Context, nc *nats.Conn) {
		<-ctx.Done()
		nc.Drain()
	}(ctx, nc)

	return api.chToModule, api.chFromModule, nil
}
