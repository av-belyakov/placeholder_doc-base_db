package natsapi

import (
	"context"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

// subscriptionHandler обработчик подписок
func (api *apiNatsModule) subscriptionHandler(ctx context.Context) {
	listener := []string{
		api.subscriptions.listenerAlert,
		api.subscriptions.listenerCase,
	}

	for _, v := range listener {
		api.natsConn.Subscribe(v, func(m *nats.Msg) {
			api.chFromModule <- SettingsOutputChan{
				TaskId:      uuid.NewString(),
				SubjectType: v,
				Data:        m.Data,
			}
		})
	}
}

// incomingInformationHandler обработчик информации полученной изнутри приложения
func (api *apiNatsModule) incomingInformationHandler(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return

		case msg := <-api.chToModule:

		}
	}
}
