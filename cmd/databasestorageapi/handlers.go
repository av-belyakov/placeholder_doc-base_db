package databasestorageapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// GetExistingIndexes проверка наличия индексов соответствующих определенному шаблону
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

// GetIndexSetting настройки выбранного индекса
func (dbs *DatabaseStorage) GetIndexSetting(ctx context.Context, index, query string) (
	settings map[string]struct {
		Settings struct {
			Index struct {
				Mapping struct {
					TotalFields struct {
						Limit string `json:"limit"`
					} `json:"total_fields"`
				} `json:"mapping"`
			} `json:"index"`
		} `json:"settings"`
	}, err error) {
	req := esapi.IndicesGetSettingsRequest{
		Index:  []string{index},
		Pretty: true,
		Human:  true,
	}

	res, err := req.Do(ctx, dbs.client.Transport)
	if err != nil {
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("the server response when executing an index search query is equal to '%s'", res.Status())

		return
	}

	err = json.NewDecoder(res.Body).Decode(&settings)
	if err != nil {
		return
	}

	return
}

// SetIndexSetting новые настройки индекса
func (dbs *DatabaseStorage) SetIndexSetting(ctx context.Context, indexes []string, query string) (bool, error) {
	indicesSettings := esapi.IndicesPutSettingsRequest{
		Index: indexes,
		Body:  strings.NewReader(query),
	}

	res, err := indicesSettings.Do(ctx, dbs.client.Transport)
	if err != nil {
		return false, err
	}
	defer responseClose(res)

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return true, nil
	}

	r := map[string]any{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		_, f, l, _ := runtime.Caller(0)
		return true, fmt.Errorf("'%v' %s:%d", err, f, l-1)
	}

	if e, ok := r["error"]; ok {
		return true, fmt.Errorf("received from module Elsaticsearch: %s (%s)", res.Status(), e)
	}

	return false, nil
}

// DelIndexSetting
func (dbs *DatabaseStorage) DelIndexSetting(ctx context.Context, indexes []string) error {
	req := esapi.IndicesDeleteRequest{Index: indexes}
	res, err := req.Do(ctx, dbs.client.Transport)
	if err != nil {
		return err
	}
	defer responseClose(res)

	return err
}

// InsertDocument добавить новый документ в заданный индекс
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
		return res, supportingfunctions.CustomError(fmt.Errorf("received from database: %s (%s), %v", tag, res.Status(), err))
	}

	return res, nil
}

// UpdateDocument поиск и обновление документов
func (dbs *DatabaseStorage) UpdateDocument(tag, currentIndex string, list []ServiseOption, document []byte) (res *esapi.Response, countDel int, err error) {
	for _, v := range list {
		res, errDel := dbs.client.Delete(v.Index, v.ID)
		responseClose(res)
		if errDel != nil {
			err = fmt.Errorf("%v, %v", err, errDel)
		}

		countDel++
	}

	res, err = dbs.InsertDocument(tag, currentIndex, document)

	return res, countDel, err
}

// SetMaxTotalFieldsLimit устанавливает максимальный лимит полей для
// переданного списка индексов в 2000, если такой лимит не был установлен ранее.
// Данная функция позволяет убрать ошибку 'Elasticsearch типа Limit of total
// fields [1000] has been exceeded while adding new fields' которая
// возникает при установленном максимальном количестве полей в 1000, что
// соответствует дефолтному значению.
func (dbs *DatabaseStorage) SetMaxTotalFieldsLimit(ctx context.Context, indexes []string) error {
	if len(indexes) == 0 {
		return fmt.Errorf("an empty list of indexes was received")
	}

	getIndexLimit := func(ctx context.Context, indexName string) (string, bool, error) {
		indexSettings, err := dbs.GetIndexSetting(ctx, indexName, "")
		if err != nil {
			return "", false, err
		}

		if info, ok := indexSettings[indexName]; ok {
			return info.Settings.Index.Mapping.TotalFields.Limit, ok, nil
		}

		return "", false, nil
	}

	var errList error
	indexForTotalFieldsLimit := []string(nil)
	for _, v := range indexes {
		limit, ok, err := getIndexLimit(ctx, v)
		if err != nil {
			errList = errors.Join(errList, supportingfunctions.CustomError(err))
		}

		if !ok || limit == "2000" {
			continue
		}

		indexForTotalFieldsLimit = append(indexForTotalFieldsLimit, v)
	}

	if len(indexForTotalFieldsLimit) == 0 {
		return errList
	}

	var query string = `{
		"index": {
			"mapping": {
				"total_fields": {
					"limit": 2000
					}
				}
			}
		}`
	if _, err := dbs.SetIndexSetting(ctx, indexForTotalFieldsLimit, query); err != nil {
		errList = errors.Join(errList, err)

		return err
	}

	return errList
}

// SearchUnderlineIdAlert поиск объекта типа 'alert' по его _id
func (dbs *DatabaseStorage) SearchUnderlineIdAlert(ctx context.Context, indexName, rootId, source string) (string, error) {
	var alertId string

	query := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", source, rootId))

	//выполняем поиск _id индекса
	res, err := dbs.client.Search(
		dbs.client.Search.WithContext(ctx),
		dbs.client.Search.WithIndex(indexName),
		dbs.client.Search.WithBody(query),
	)
	if err != nil {
		return alertId, err
	}
	defer responseClose(res)

	if res.StatusCode != http.StatusOK {
		return alertId, fmt.Errorf("%s", res.Status())
	}

	tmp := AlertDBResponse{}
	if err = json.NewDecoder(res.Body).Decode(&tmp); err != nil {
		return alertId, err
	}

	for _, v := range tmp.Options.Hits {
		alertId = v.ID
	}

	return alertId, nil
}

// SearchUnderlineIdCase поиск объекта типа 'case' по его _id
func (dbs *DatabaseStorage) SearchUnderlineIdCase(ctx context.Context, indexName, rootId, source string) (string, error) {
	var caseId string
	query := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", source, rootId))

	//выполняем поиск _id индекса
	res, err := dbs.client.Search(
		dbs.client.Search.WithContext(ctx),
		dbs.client.Search.WithIndex(indexName),
		dbs.client.Search.WithBody(query),
	)
	if err != nil {
		return caseId, err
	}
	defer responseClose(res)

	if res.StatusCode != http.StatusOK {
		return caseId, fmt.Errorf("%s", res.Status())
	}

	tmp := CaseDBResponse{}
	if err = json.NewDecoder(res.Body).Decode(&tmp); err != nil {
		return caseId, err
	}

	for _, v := range tmp.Options.Hits {
		caseId = v.ID
	}

	return caseId, nil
}
