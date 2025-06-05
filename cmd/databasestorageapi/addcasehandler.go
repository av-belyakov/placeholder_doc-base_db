package databasestorageapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/request"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// addCase добавление объекта типа 'case'
func (dbs *DatabaseStorage) addCase(ctx context.Context, data any) {
	fmt.Println("func 'DatabaseStorage.addCase' START")

	t := time.Now()

	newDocument, ok := data.(*documentgenerator.VerifiedCase)
	if !ok {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("type conversion error")).Error())

		return
	}

	//получаем наименование хранилища
	indexName, isExist := dbs.settings.storages["case"]
	if !isExist {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	//формируем наименование индекса
	currentIndex := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), int(t.Month()))

	caseId := fmt.Sprint(newDocument.GetEvent().GetObject().CaseId)
	reqSetTag := fmt.Appendf(nil, `{
						  "service": "placeholder_doc-base_db",
						  "command": "add_case_tag",
						  "root_id": "%s",
						  "case_id": "%s",
						  "value": "Webhook: send=\"ElasticsearchDB"
						}`,
		newDocument.GetEvent().GetRootId(),
		caseId)

	//получаем список ip адресов
	ipAddrObjects := newDocument.GetAdditionalInformation().GetIpAddressesInformation()
	reqGeoIP, err := json.Marshal(request.RequestGeoIP{
		Source:          "placeholder_doc-base_db",
		TaskId:          newDocument.GetEvent().GetRootId(),
		ListIpAddresses: documentgenerator.GetListIPAddr(ipAddrObjects),
	})
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
	}

	//получаем список сенсоров
	sensorIdObjects := newDocument.GetAdditionalInformation().GetSensorsInformation()
	reqSensorId, err := json.Marshal(request.RequestSensorInformation{
		Source:      "placeholder_doc-base_db",
		TaskId:      newDocument.GetEvent().GetRootId(),
		ListSensors: documentgenerator.GetListSensorId(sensorIdObjects),
	})

	newDocumentBinary, err := json.Marshal(newDocument.Get())
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//получаем существующие индексы
	existingIndexes, err := dbs.GetExistingIndexes(ctx, indexName)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//будет выполнятся поиск по индексам только в текущем году так как при
	//накоплении большого количества индексов, поиск по всем серьезно замедлит работу
	indexesOnlyCurrentYear := []string(nil)
	for _, v := range existingIndexes {
		if strings.Contains(v, fmt.Sprint(t.Year())) {
			indexesOnlyCurrentYear = append(indexesOnlyCurrentYear, v)
		}
	}

	// если похожих индексов нет
	if len(indexesOnlyCurrentYear) == 0 {
		//
		//вставка документа
		statusCode, err := dbs.InsertDocument(ctx, currentIndex, newDocumentBinary)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}

		existingIndexes = append(existingIndexes, currentIndex)
		//устанавливаем максимальный лимит количества полей для всех индексов которые
		//содержат значение по умолчанию в 1000 полей
		if err := dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes); err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
		}

		//счетчик
		dbs.counter.SendMessage("update count insert subject case to db", 1)
		dbs.logger.Send("info", fmt.Sprintf("insert new document to caseId:'%s' with rootId:'%s', status code:'%d'", caseId, newDocument.GetEvent().GetRootId(), statusCode))

		//запрос на отправку сообщения для установки тега
		dbs.GetChanDataFromModule() <- SettingsChanOutput{
			Command: "set_tag",
			CaseId:  caseId,
			RootId:  newDocument.GetEvent().GetRootId(),
			Data:    reqSetTag,
		}

		if len(ipAddrObjects) > 0 {
			//запросы для обогащения кейса дополнительной информацией геопозиционирование
			// ip адресов
			dbs.GetChanDataFromModule() <- SettingsChanOutput{
				Command: "get_geoip_info",
				RootId:  newDocument.GetEvent().GetRootId(),
				Data:    reqGeoIP,
			}
		}

		if len(sensorIdObjects) > 0 {
			//запросы для обогащения кейса дополнительной информацией информация о
			// принадлежности и месте установки сенсора
			dbs.GetChanDataFromModule() <- SettingsChanOutput{
				Command: "get_sensor_info",
				RootId:  newDocument.GetEvent().GetRootId(),
				Data:    reqSensorId,
			}
		}

		return
	}

	//устанавливаем максимальный лимит количества полей для всех индексов которые
	//содержат значение по умолчанию в 1000 полей
	if err := dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
	}

	currentQuery := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", newDocument.GetSource(), newDocument.GetEvent().GetRootId()))
	res, err := dbs.client.Search(
		dbs.client.Search.WithContext(context.Background()),
		dbs.client.Search.WithIndex(indexesOnlyCurrentYear...),
		dbs.client.Search.WithBody(currentQuery),
	)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}
	defer res.Body.Close()

	response := CaseDBResponse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//вставка выполняется только когда не найден искомый документ
	if response.Options.Total.Value == 0 {
		//
		//вставка документа
		statusCode, err := dbs.InsertDocument(ctx, currentIndex, newDocumentBinary)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}

		//счетчик
		dbs.counter.SendMessage("update count insert subject case to db", 1)
		dbs.logger.Send("info", fmt.Sprintf("insert new document to caseId:'%s' with rootId:'%s', status code:'%d'", caseId, newDocument.GetEvent().GetRootId(), statusCode))

		//запрос на отправку сообщения для установки тега
		dbs.GetChanDataFromModule() <- SettingsChanOutput{
			Command: "set_tag",
			CaseId:  caseId,
			RootId:  newDocument.GetEvent().GetRootId(),
			Data:    reqSetTag,
		}

		if len(ipAddrObjects) > 0 {
			//запросы для обогащения кейса дополнительной информацией геопозиционирование
			// ip адресов
			dbs.GetChanDataFromModule() <- SettingsChanOutput{
				Command: "get_geoip_info",
				RootId:  newDocument.GetEvent().GetRootId(),
				Data:    reqGeoIP,
			}
		}

		if len(sensorIdObjects) > 0 {
			//запросы для обогащения кейса дополнительной информацией информация о
			// принадлежности и месте установки сенсора
			dbs.GetChanDataFromModule() <- SettingsChanOutput{
				Command: "get_sensor_info",
				RootId:  newDocument.GetEvent().GetRootId(),
				Data:    reqSensorId,
			}
		}

		return
	}

	//*** при наличие искомого документа выполняем его замену ***
	//***********************************************************
	var countReplacingFields int
	listDeleting := []ServiseOption(nil)
	updateVerified := documentgenerator.NewVerifiedCase()
	for _, v := range response.Options.Hits {
		count, err := updateVerified.Event.ReplacingOldValues(*v.Source.GetEvent())
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
		} else {
			countReplacingFields += count
		}

		countReplacingFields += updateVerified.GetObservables().ReplacingOldValues(v.Source.Observables)
		countReplacingFields += updateVerified.GetTtps().ReplacingOldValues(v.Source.Ttps)

		listDeleting = append(listDeleting, ServiseOption{
			ID:    v.ID,
			Index: v.Index,
		})

		//устанавливаем время создания первой записи о кейсе
		updateVerified.SetCreateTimestamp(v.Source.CreateTimestamp)
	}

	//выполняем обновление объекта типа Event
	updateVerified.SetSource(newDocument.GetSource())
	num, err := updateVerified.Event.ReplacingOldValues(*newDocument.GetEvent())
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
	} else {
		countReplacingFields += num
	}

	countReplacingFields += updateVerified.GetObservables().ReplacingOldValues(*newDocument.GetObservables())
	countReplacingFields += updateVerified.GetTtps().ReplacingOldValues(*newDocument.GetTtps())

	nvbyte, err := json.Marshal(updateVerified)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//обновление уже существующего документа
	statusCode, countDel, err := dbs.UpdateDocument(ctx, currentIndex, listDeleting, nvbyte)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("rootId '%s' '%s'", newDocument.GetEvent().GetRootId(), err.Error())).Error())

		return
	}

	dbs.counter.SendMessage("update count insert subject case to db", 1)
	dbs.logger.Send("info", fmt.Sprintf("update document 'case' type, count delete:'%d', count replacing fields:'%d' for case with rootId:'%s', status code:'%d'", countDel, countReplacingFields, newDocument.GetEvent().GetRootId(), statusCode))
}
