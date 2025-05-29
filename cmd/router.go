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
				switch msg.SubjectType {
				case "alert":
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

				case "case":
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
					}()

				case "geoip information":
					//передача информации о географическом местоположении ip адресов
					r.chToDBSApi <- databasestorageapi.SettingsChanInput{
						Section: "information handling",
						Command: "add geoip information",
						Data:    msg.Data,
					}

				case "sensor information":
					// местоположении и принадлежности сенсоров
					r.chToDBSApi <- databasestorageapi.SettingsChanInput{
						Section: "information handling",
						Command: "add sensor information",
						Data:    msg.Data,
					}

				default:
					r.logger.Send("error", supportingfunctions.CustomError(errors.New("undefined subscription type")).Error())

				}

			case msg := <-r.chFromDBSApi:
				//пересылаются запросы на установку тега в TheHive, geoip информации
				// информации о место располажения и принадлежности сенсоров
				r.chToNatsApi <- natsapi.SettingsChanInput{
					Command: msg.Command,
					RootId:  msg.RootId,
					Data:    msg.Data,
				}
			}
		}
	}()
}
