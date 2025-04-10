package main

import (
	"context"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/countermessage"
)

// NewRouter маршрутизатор сообщений внутри приложения
func NewRouter(counter *countermessage.CounterMessage, logger interfaces.Logger, settings ApplicationRouterSettings) *ApplicationRouter {
	return &ApplicationRouter{
		counter:       counter,
		logger:        logger,
		chToNatsApi:   settings.ChanToNats,
		chFromNatsApi: settings.ChanFromNats,
		chToDBSApi:    settings.ChanToDBS,
		chFromDBSApi:  settings.ChanFromDBS,
	}
}

func (r *ApplicationRouter) Router(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():

			case msg := <-r.chFromNatsApi:
				switch msg.SubjectType {
				case "object.alerttype":
					go func() {
						decoder := decoderjsondocuments.New(r.counter, r.logger)
						chDecode := decoder.Start(msg.Data, msg.TaskId)

						rootId, verifyAlert, listRawFields := documentgenerator.AlertGenerator(chDecode)
					}()

				case "object.casetype":
					go func() {
						decoder := decoderjsondocuments.New(r.counter, r.logger)
						chDecode := decoder.Start(msg.Data, msg.TaskId)

						rootId, verifyCase, listRawFields := documentgenerator.CaseGenerator(chDecode)
						/*
							//отправляем запрос в модуль NATS для установки тега 'Webhook: send="ES"'
							opts.natsChan <- natsinteractions.SettingsInputChan{
								Command: "send tag",
								EventId: fmt.Sprint(objectElem.GetCaseId()),
								TaskId:  opts.msgId,
							}

							//делаем запрос на получение дополнительной информации о сенсорах
							if len(sensorsId.Get()) > 0 || len(ipAddresses.Get()) > 0 {
								//делаем запрос к модулю обогащения доп. информацией из Zabbix
								opts.eemChan <- eventenrichmentmodule.SettingsChanInputEEM{
									RootId:      eventCase.GetRootId(),
									Source:      verifiedCase.GetSource(),
									SensorsId:   sensorsId.Get(),
									IpAddresses: ipAddresses.Get(),
								}
							}
						*/
					}()

				}

			case msg := <-r.chFromDBSApi:

			}
		}
	}()
}
