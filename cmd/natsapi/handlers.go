package natsapi

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// subscriptionHandler обработчик подписок
func (api *apiNatsModule) subscriptionHandler() {
	for _, v := range api.subscriptions {
		_, err := api.natsConn.Subscribe(v, func(m *nats.Msg) {
			api.chFromModule <- SettingsChanOutput{
				TaskId:      uuid.NewString(),
				SubjectType: v,
				Data:        m.Data,
			}
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
			//команда на установку тега
			if err := api.natsConn.Publish(api.settings.command,
				fmt.Appendf(nil, `{
									  "service": "placeholder_docbase_db",
									  "command": "add_case_tag",
									  "root_id": "%s",
									  "case_id": "%s",
									  "value": "Webhook: send=\"ES\""
								}`, incomingData.RootId, incomingData.CaseId)); err != nil {
				api.logger.Send("error", supportingfunctions.CustomError(err).Error())
			}
		}
	}
}
