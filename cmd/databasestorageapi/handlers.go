package databasestorageapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8/esapi"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/documentgenerator"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

func (dbs *DatabaseStorage) addAlert(ctx context.Context, data any) {
	newDocument, ok := data.(documentgenerator.VerifiedAlert)
	if !ok {
		dbs.logging.Send("error", supportingfunctions.CustomError(errors.New("type conversion error")).Error())

		return
	}

}

func (dbs *DatabaseStorage) addCase(ctx context.Context, data any) {
	newDocument, ok := data.(documentgenerator.VerifiedCase)
	if !ok {
		dbs.logging.Send("error", supportingfunctions.CustomError(errors.New("type conversion error")).Error())

		return
	}

	indexName, isExist := dbs.settings.storages["case"]
	if !isExist {
		dbs.logging.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	var countReplacingFields int
	tag := fmt.Sprintf("case rootId: '%s'", newDocument.GetEvent().GetRootId())

	t := time.Now()
	month := int(t.Month())
	indexCurrent := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), month)
	queryCurrent := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", newDocument.GetSource(), newDocument.GetEvent().GetRootId()))

	sensorsId := listSensorId{
		sensors: []string(nil),
	}

	caseElem := newDocument.Get()
	eventElem := caseElem.GetEvent()
	objectElem := eventElem.GetObject()
	if listSensorId, ok := objectElem.GetTags()["sensor:id"]; ok {
		for _, v := range listSensorId {
			sensorsId.addElem(v)
		}
	}

	newDocumentBinary, err := json.Marshal(newDocument.Get())
	if err != nil {
		dbs.logging.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	indexes, err := dbs.GetExistingIndexes(ctx, indexName)
	if err != nil {
		dbs.logging.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	//будет выполнятся поиск по индексам только в текущем году
	//так как при накоплении большого количества индексов, поиск
	//по всем серьезно замедлит работу
	indexesOnlyCurrentYear := []string(nil)
	for _, v := range indexes {
		if strings.Contains(v, fmt.Sprint(t.Year())) {
			indexesOnlyCurrentYear = append(indexesOnlyCurrentYear, v)
		}
	}

	// если похожих индексов нет
	if len(indexesOnlyCurrentYear) == 0 {
		hsd.InsertNewDocument(tag, indexCurrent, newDocumentBinary, logging, counting)

		/*
		   res, err := hsd.InsertDocument(tag, index, document)
		   	defer responseClose(res)
		   	if err != nil {
		   		_, f, l, _ := runtime.Caller(0)
		   		logging <- datamodels.MessageLogging{
		   			MsgData: fmt.Sprintf("'%s' %s:%d", err.Error(), f, l-2),
		   			MsgType: "error",
		   		}

		   		return
		   	}

		   	//счетчик
		   	counting <- datamodels.DataCounterSettings{
		   		DataType: "update count insert Elasticserach",
		   		DataMsg:  "subject_alert",
		   		Count:    1,
		   	}
		*/

		indexes = append(indexes, indexCurrent)
		//устанавливаем максимальный лимит количества полей для всех индексов которые
		//содержат значение по умолчанию в 1000 полей
		if err := SetMaxTotalFieldsLimit(hsd, indexes, logging); err != nil {
			dbs.logging.Send("error", supportingfunctions.CustomError(err).Error())

			return
		}
	}

	//устанавливаем максимальный лимит количества полей для всех индексов которые
	//содержат значение по умолчанию в 1000 полей
	if err := SetMaxTotalFieldsLimit(hsd, indexes, logging); err != nil {
		dbs.logging.Send("error", supportingfunctions.CustomError(err).Error())
	}

	res, err := hsd.SearchDocument(indexesOnlyCurrentYear, queryCurrent)
	defer func() {
		if res == nil || res.Body == nil {
			return
		}

		res.Body.Close()
	}()
	if err != nil {
		dbs.logging.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	decEs := datamodels.ElasticsearchResponseCase{}
	err = json.NewDecoder(res.Body).Decode(&decEs)
	if err != nil {
		dbs.logging.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	if decEs.Options.Total.Value == 0 {
		//выполняется только когда не найден искомый документ
		hsd.InsertNewDocument(tag, indexCurrent, newDocumentBinary, logging, counting)

		return
	}

	//*** при наличие искомого документа выполняем его замену ***
	//***********************************************************
	listDeleting := []datamodels.ServiseOption(nil)
	updateVerified := datamodels.NewVerifiedEsCase()
	for _, v := range decEs.Options.Hits {
		count, err := updateVerified.Event.ReplacingOldValues(*v.Source.GetEvent())
		if err != nil {
			dbs.logging.Send("error", supportingfunctions.CustomError(err).Error())
		} else {
			countReplacingFields += count
		}

		countReplacingFields += updateVerified.ObservablesMessageEs.ReplacingOldValues(v.Source.ObservablesMessageEs)
		countReplacingFields += updateVerified.TtpsMessageTheHive.ReplacingOldValues(v.Source.TtpsMessageTheHive)

		listDeleting = append(listDeleting, datamodels.ServiseOption{
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
		dbs.logging.Send("error", supportingfunctions.CustomError(err).Error())
	} else {
		countReplacingFields += num
	}

	countReplacingFields += updateVerified.ObservablesMessageEs.ReplacingOldValues(*newDocument.GetObservables())
	countReplacingFields += updateVerified.TtpsMessageTheHive.ReplacingOldValues(*newDocument.GetTtps())

	nvbyte, err := json.Marshal(updateVerified)
	if err != nil {
		dbs.logging.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	res, countDel, err := hsd.UpdateDocument(tag, indexCurrent, listDeleting, nvbyte)
	defer responseClose(res)
	if err != nil {
		dbs.logging.Send("error", supportingfunctions.CustomError(fmt.Errorf("rootId '%s' '%s'", newDocument.GetEvent().GetRootId(), err.Error())).Error())

		return
	}

	if res != nil && res.StatusCode == http.StatusCreated {
		//счетчик
		counting <- datamodels.DataCounterSettings{
			DataType: "update count insert Elasticserach",
			DataMsg:  "subject_case",
			Count:    1,
		}

		dbs.logging.Send("warning", supportingfunctions.CustomError(fmt.Errorf("count delete: '%d', count replacing fields '%d' for alert with rootId: '%s'", countDel, countReplacingFields, newDocument.GetEvent().GetRootId())).Error())
	}
}

// GetExistingIndexes выполняет проверку наличия индексов соответствующих определенному
// шаблону и возвращает список наименований индексов подходящих под заданный шаблон
func (dbs *DatabaseStorage) GetExistingIndexes(ctx context.Context, pattern string) ([]string, error) {
	listIndexes := []string(nil)
	msg := []struct {
		Index string `json:"index"`
	}(nil)

	res, err := dbs.client.Cat.Indices(
		dbs.client.Cat.Indices.WithContext(ctx),
		dbs.client.Cat.Indices.WithFormat("json"),
	)
	if err != nil {
		return nil, err
	}
	defer responseClose(res)

	if err = json.NewDecoder(res.Body).Decode(&msg); err != nil {
		return nil, err
	}

	for _, v := range msg {
		if !strings.Contains(v.Index, pattern) {
			continue
		}

		listIndexes = append(listIndexes, v.Index)
	}

	return listIndexes, err
}

// GetIndexSetting получает натройки выбранного индекса
func (dbs *DatabaseStorage) GetIndexSetting(index, query string) (*esapi.Response, error) {
	var (
		res *esapi.Response
		err error
	)

	req := esapi.IndicesGetSettingsRequest{
		Index:  []string{index},
		Pretty: true,
		Human:  true,
	}

	res, err = req.Do(context.Background(), dbs.client.Transport)
	if err != nil {
		return res, err
	}

	return res, nil
}

// InsertDocument добавляет новый документ в заданный индекс
func (dbs *DatabaseStorage) InsertDocument(tag, index string, b []byte) (*esapi.Response, error) {
	var res *esapi.Response

	if dbs.client == nil {
		return res, supportingfunctions.CustomError(errors.New("the client parameters for connecting to the Elasticsearch database are not set correctly"))
	}

	buf := bytes.NewReader(b)
	res, err := dbs.client.Index(index, buf)
	defer responseClose(res)
	if err != nil {
		return res, supportingfunctions.CustomError(err)
	}

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return res, nil
	}

	r := map[string]any{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return res, supportingfunctions.CustomError(err)
	}

	if err, ok := r["error"]; ok {
		return res, supportingfunctions.CustomError(fmt.Errorf("%s received from module Elsaticsearch: %s (%s), %w", tag, res.Status(), err))
	}

	return res, nil
}
