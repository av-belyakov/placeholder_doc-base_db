package main

import (
	"context"
	"errors"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/databasestorageapi"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/decoderjsondocuments"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-base_db/cmd/natsapi"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// NewRouter маршрутизатор сообщений внутри приложения
func NewRouter(counter interfaces.Counter, logger interfaces.Logger, settings ApplicationRouterSettings) *ApplicationRouter {
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
		decoder := decoderjsondocuments.New(r.counter, r.logger)

		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-r.chFromNatsApi:
				if msg.SubjectType == "alert" {
					go func() {
						rootId, verifyAlert, listRawFields := documentgenerator.AlertGenerator(decoder.Start(msg.Data, msg.TaskId))

						if len(listRawFields) > 0 {
							r.logger.Send("alert_raw_fields", supportingfunctions.JoinRawFieldsToString(listRawFields, "rootId", rootId))
						}

						r.chToDBSApi <- databasestorageapi.SettingsChanInput{
							Section: "handling alert",
							Command: "add alert",
							Data:    verifyAlert,
						}
					}()
				} else if msg.SubjectType == "case" {
					go func() {
						rootId, verifyCase, listRawFields := documentgenerator.CaseGenerator(decoder.Start(msg.Data, msg.TaskId))

						if len(listRawFields) > 0 {
							r.logger.Send("case_raw_fields", supportingfunctions.JoinRawFieldsToString(listRawFields, "rootId", rootId))
						}

						//передача объекта в модуль взаимодействия с базой данных для
						//дальнейшей загрузки данных в базу
						r.chToDBSApi <- databasestorageapi.SettingsChanInput{
							Section: "handling case",
							Command: "add case",
							Data:    verifyCase,
						}

						ipAddrObjects := verifyCase.GetAdditionalInformation().GetIpAddressesInformation()

						//запрос на получение дополнительной информации об ip адресе (GeoIP)
						listIpAddr := documentgenerator.GetListIPAddr(ipAddrObjects)

						r.chToNatsApi <- natsapi.SettingsChanInput{}

						//запрос на получение дополнительной информации о сенсоре

					}()
				} else {
					r.logger.Send("error", supportingfunctions.CustomError(errors.New("undefined subscription type")).Error())
				}

			case msg := <-r.chFromDBSApi:
				//запрос на установку тега в TheHive
				r.chToNatsApi <- natsapi.SettingsChanInput{
					Command: msg.Command,
					RootId:  msg.RootId,
				}

				//
				//
				// Здесь же надо отправить запрос, через NATS, к сторониим модулям
				// для получения дополнительной информации, такой как GeoIP, подробная
				// информация о сенсорах и т.д.
				//
				//

				/*
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

			}
		}
	}()
}
