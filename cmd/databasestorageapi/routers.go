package databasestorageapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

func (dbs *DatabaseStorage) router(ctx context.Context) {
	handlersList := map[string]map[string]func(context.Context, any){
		"handling alert": {
			"add alert": dbs.addAlert,
		},
		"handling case": {
			"add case": dbs.addCase,
		},
	}

	for {
		select {
		case <-ctx.Done():
			return

		case msg := <-dbs.chInput:
			strErr := "the handler for the accepted request was not found"

			if _, ok := handlersList[msg.Section]; !ok {
				dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(strErr)).Error())

				continue
			}

			if _, ok := handlersList[msg.Section][msg.Command]; !ok {
				dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(strErr)).Error())

				continue
			}

			go handlersList[msg.Section][msg.Command](ctx, msg.Data)
		}
	}
}

// addAlert добавление объекта типа 'alert'
func (dbs *DatabaseStorage) addAlert(ctx context.Context, data any) {
	/*newDocument, ok := data.(documentgenerator.VerifiedAlert)
	if !ok {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("type conversion error")).Error())

		return
	}*/
}

// addCase добавление объекта типа 'case'
func (dbs *DatabaseStorage) addCase(ctx context.Context, data any) {
	newDocument, ok := data.(documentgenerator.VerifiedCase)
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

	t := time.Now()
	//формируем наименование индекса
	currentIndex := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), int(t.Month()))
	currentQuery := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", newDocument.GetSource(), newDocument.GetEvent().GetRootId()))

	sensorsId := listSensorId{
		sensors: []string(nil),
	}

	objectElement := newDocument.Get().GetEvent().GetObject()
	tag := fmt.Sprintf("case rootId: '%s'", newDocument.GetEvent().GetRootId())
	if listSensorId, ok := objectElement.GetTags()["sensor:id"]; ok {
		for _, v := range listSensorId {
			sensorsId.addElem(v)
		}
	}

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
		res, err := dbs.InsertDocument(tag, currentIndex, newDocumentBinary)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}
		defer responseClose(res)

		//счетчик
		dbs.counter.SendMessage("update count insert subject case to db", 1)

		existingIndexes = append(existingIndexes, currentIndex)
		//устанавливаем максимальный лимит количества полей для всех индексов которые
		//содержат значение по умолчанию в 1000 полей
		if err := dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes); err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
		}

		return
	}

	//устанавливаем максимальный лимит количества полей для всех индексов которые
	//содержат значение по умолчанию в 1000 полей
	if err := dbs.SetMaxTotalFieldsLimit(ctx, existingIndexes); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())
	}

	res, err := dbs.client.Search(
		dbs.client.Search.WithContext(context.Background()),
		dbs.client.Search.WithIndex(indexesOnlyCurrentYear...),
		dbs.client.Search.WithBody(currentQuery),
	)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}
	defer responseClose(res)

	response := CaseDBResponse{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//вставка выполняется только когда не найден искомый документ
	if response.Options.Total.Value == 0 {
		res, err := dbs.InsertDocument(tag, currentIndex, newDocumentBinary)
		if err != nil {
			dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}
		defer responseClose(res)

		//счетчик
		dbs.counter.SendMessage("update count insert subject case to db", 1)

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

	res, countDel, err := dbs.UpdateDocument(tag, currentIndex, listDeleting, nvbyte)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("rootId '%s' '%s'", newDocument.GetEvent().GetRootId(), err.Error())).Error())

		return
	}
	defer responseClose(res)

	if res != nil && res.StatusCode == http.StatusCreated {
		dbs.counter.SendMessage("update count insert subject case to db", 1)
		dbs.logger.Send("warning", supportingfunctions.CustomError(fmt.Errorf("count delete: '%d', count replacing fields '%d' for alert with rootId: '%s'", countDel, countReplacingFields, newDocument.GetEvent().GetRootId())).Error())
	}
}
