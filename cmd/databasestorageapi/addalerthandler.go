package databasestorageapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// addAlert добавление объекта типа 'alert'
func (dbs *DatabaseStorage) addAlert(ctx context.Context, data any) {
	t := time.Now()
	ctxTimeout, ctxCancel := context.WithTimeout(ctx, time.Second*15)
	defer ctxCancel()

	newDocument, ok := data.(*documentgenerator.VerifiedAlert)
	if !ok {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("type conversion error")).Error())

		return
	}

	//получаем наименование хранилища
	indexName, isExist := dbs.settings.storages["alert"]
	if !isExist {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	//формируем наименование индекса
	currentIndex := fmt.Sprintf("%s_%s_%d_%d", indexName, newDocument.GetSource(), t.Year(), int(t.Month()))

	newDocumentBinary, err := json.Marshal(newDocument.Get())
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//получаем существующие индексы
	existingIndexes, err := dbs.GetExistingIndexes(ctxTimeout, indexName)
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
		statusCode, err := dbs.InsertDocument(ctxTimeout, currentIndex, newDocumentBinary)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}

		//устанавливаем максимальный лимит количества полей для всех индексов которые
		//содержат значение по умолчанию в 1000 полей
		if err := dbs.SetMaxTotalFieldsLimit(ctxTimeout, existingIndexes); err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
		}

		//счетчик
		dbs.counter.SendMessage("update count insert subject alert to db", 1)
		dbs.logger.Send("info", fmt.Sprintf("insert new document 'alert' type, with rootId:'%s', status code:'%d'", newDocument.GetEvent().GetRootId(), statusCode))

		return
	}

	//устанавливаем максимальный лимит количества полей для всех индексов которые
	//содержат значение по умолчанию в 1000 полей
	if err := dbs.SetMaxTotalFieldsLimit(ctxTimeout, existingIndexes); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
	}

	//ищем alert с таким же rootId
	res, err := dbs.client.Search(
		dbs.client.Search.WithContext(ctxTimeout),
		dbs.client.Search.WithIndex(indexesOnlyCurrentYear...),
		dbs.client.Search.WithBody(
			strings.NewReader(
				fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", newDocument.GetSource(), newDocument.GetEvent().GetRootId()),
			),
		),
	)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}
	defer res.Body.Close()

	response := AlertDBResponse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//вставка выполняется только когда не найден искомый документ
	if response.Options.Total.Value == 0 {
		//
		//вставка документа
		statusCode, err := dbs.InsertDocument(ctxTimeout, currentIndex, newDocumentBinary)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}

		//счетчик
		dbs.counter.SendMessage("update count insert subject case to db", 1)
		dbs.logger.Send("info", fmt.Sprintf("insert new document 'alert' type, with rootId:'%s', status code:'%d'", newDocument.GetEvent().GetRootId(), statusCode))

		return
	}

	//*** при наличие искомого документа выполняем его замену ***
	//***********************************************************
	var countReplacingFields int
	listDeleting := []ServiseOption(nil)
	updateVerified := documentgenerator.NewVerifiedAlert()
	for _, v := range response.Options.Hits {
		count, err := updateVerified.Event.ReplacingOldValues(*v.Source.GetEvent())
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
		} else {
			countReplacingFields += count
		}

		count, err = updateVerified.Alert.ReplacingOldValues(*v.Source.GetAlert())
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
		} else {
			countReplacingFields += count
		}

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

	//выполняем обновление объекта типа Alert
	num, err = updateVerified.Alert.ReplacingOldValues(*newDocument.GetAlert())
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
	} else {
		countReplacingFields += num
	}

	nvbyte, err := json.Marshal(updateVerified)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//обновление уже существующего документа
	statusCode, countDel, err := dbs.UpdateDocument(ctxTimeout, currentIndex, listDeleting, nvbyte)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("rootId '%s' '%s'", newDocument.GetEvent().GetRootId(), err.Error())).Error())

		return
	}

	dbs.counter.SendMessage("update count insert subject case to db", 1)
	dbs.logger.Send("info", fmt.Sprintf("update document 'alert' type, count delete:'%d', count replacing fields:'%d' for alert with rootId:'%s', status code:'%d'", countDel, countReplacingFields, newDocument.GetEvent().GetRootId(), statusCode))
}
