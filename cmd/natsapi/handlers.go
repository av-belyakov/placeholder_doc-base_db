package natsapi

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// subscriptionHandler обработчик подписок
func (api *apiNatsModule) subscriptionHandler() {
	for k, v := range api.subscriptions {
		_, err := api.natsConn.Subscribe(v, func(m *nats.Msg) {
			api.chFromModule <- SettingsChanOutput{
				TaskId:      uuid.NewString(),
				SubjectType: k,
				Data:        m.Data,
			}

			//счетчик принятых событий
			api.counter.SendMessage("update accepted events", 1)
		})
		if err != nil {
			api.logger.Send("error", supportingfunctions.CustomError(err).Error())
		}
	}
}

// incomingInformationHandler обработчик информации полученной изнутри приложения
func (api *apiNatsModule) incomingInformationHandler(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case incomingData := <-api.chToModule:
			switch incomingData.Command {
			case "set_tag":
				//команда на установку тега
				if err := api.natsConn.Publish(api.settings.command, incomingData.Data); err != nil {
					api.logger.Send("error", supportingfunctions.CustomError(err).Error())
				}

			case "get_geoip_info":
				go func(ctx context.Context) {
					ctxTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
					defer cancel()

					res, err := api.natsConn.RequestWithContext(ctxTimeout, api.requests["get_geoip_info"], incomingData.Data)
					if err != nil {
						api.logger.Send("error", supportingfunctions.CustomError(err).Error())
					}

					if res == nil {
						return
					}

					api.chFromModule <- SettingsChanOutput{
						SubjectType: "geoip information",
						Data:        res.Data,
					}
				}(ctx)

			case "get_sensor_info":
				go func(ctx context.Context) {
					ctxTimeout, cancel := context.WithTimeout(ctx, 15*time.Second)
					defer cancel()

					res, err := api.natsConn.RequestWithContext(ctxTimeout, api.requests["get_sensor_info"], incomingData.Data)
					if err != nil {
						api.logger.Send("error", supportingfunctions.CustomError(err).Error())
					}

					if res == nil {
						return
					}

					api.chFromModule <- SettingsChanOutput{
						SubjectType: "sensor information",
						Data:        res.Data,
					}
				}(ctx)

			}
		}
	}
}
