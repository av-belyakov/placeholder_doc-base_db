package databasestorageapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/goforj/godump"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// GetExistingIndexes проверка наличия индексов соответствующих определенному шаблону
func (dbs *DatabaseStorage) GetExistingIndexes(ctx context.Context, pattern string) ([]string, error) {
	if dbs.client == nil {
		return []string{}, supportingfunctions.CustomError(errors.New("the client parameters for connecting to the Elasticsearch database are not set correctly"))
	}

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
	defer res.Body.Close()

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
func (dbs *DatabaseStorage) GetIndexSetting(ctx context.Context, index string) (
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
	defer res.Body.Close()

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
	defer res.Body.Close()

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

// DelIndexSetting удаление индекса и его настроек
func (dbs *DatabaseStorage) DelIndexSetting(ctx context.Context, indexes []string) error {
	req := esapi.IndicesDeleteRequest{Index: indexes}
	res, err := req.Do(ctx, dbs.client.Transport)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return err
}

// InsertDocument добавить новый документ в заданный индекс
func (dbs *DatabaseStorage) InsertDocument(ctx context.Context, index string, b []byte) (int, error) {
	if dbs.client == nil {
		return 0, supportingfunctions.CustomError(errors.New("the client parameters for connecting to the Elasticsearch database are not set correctly"))
	}

	buf := bytes.NewReader(b)
	res, err := dbs.client.Index(index, buf)
	if err != nil {
		return 0, supportingfunctions.CustomError(err)
	}
	defer res.Body.Close()

	bodyRes, err := io.ReadAll(res.Body)
	if err != nil {
		return res.StatusCode, err
	}

	result, err := supportingfunctions.GetElementsFromJSON(ctx, bodyRes)
	if err != nil {
		return res.StatusCode, err
	}

	for k, v := range result {
		if strings.Contains(k, "error") {
			return res.StatusCode, errors.New(fmt.Sprint(v.Value))
		}
	}

	return res.StatusCode, nil
}

// UpdateDocument поиск и обновление документов
func (dbs *DatabaseStorage) UpdateDocument(ctx context.Context, currentIndex string, list []ServiseOption, document []byte) (statusCode, countDel int, err error) {
	for _, v := range list {
		res, errDel := dbs.client.Delete(v.Index, v.ID)
		if errDel != nil {
			err = fmt.Errorf("%v, %v", err, errDel)
		}
		res.Body.Close()

		countDel++
	}

	statusCode, err = dbs.InsertDocument(ctx, currentIndex, document)

	return statusCode, countDel, err
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
		indexSettings, err := dbs.GetIndexSetting(ctx, indexName)
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

// SearchUnderlineIdAlert поиск объекта типа 'alert' по его rootId
// возвращает _id объекта под которым он находится в БД
func (dbs *DatabaseStorage) SearchUnderlineIdAlert(ctx context.Context, indexName, rootId, source string) (string, error) {
	query := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"source\": \"%s\"}}, {\"match\": {\"event.rootId\": \"%s\"}}]}}}", source, rootId))

	//выполняем поиск _id индекса
	res, err := dbs.client.Search(
		dbs.client.Search.WithContext(ctx),
		dbs.client.Search.WithIndex(indexName),
		dbs.client.Search.WithBody(query),
	)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("%s", res.Status())
	}

	tmp := AlertDBResponse{}
	if err = json.NewDecoder(res.Body).Decode(&tmp); err != nil {
		return "", err
	}

	for _, v := range tmp.Options.Hits {
		return v.ID, nil
	}

	return "", nil
}

// SearchUnderlineIdCase поиск объекта типа 'case' по его rootId
// возвращает _id объекта под которым он находится в БД
func (dbs *DatabaseStorage) SearchUnderlineIdCase(ctx context.Context, indexName, rootId string) (string, error) {
	query := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"event.rootId\": \"%s\"}}]}}}", rootId))

	//выполняем поиск _id индекса
	res, err := dbs.client.Search(
		dbs.client.Search.WithContext(ctx),
		dbs.client.Search.WithIndex(indexName),
		dbs.client.Search.WithBody(query),
	)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	bodyRes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	result, err := supportingfunctions.GetElementsFromJSON(ctx, bodyRes)
	if err != nil {
		return "", err
	}

	fmt.Println("func 'SearchUnderlineIdCase'")
	godump.Dump(result)

	for k, v := range result {
		if k == "hits.hits._id" {
			return fmt.Sprint(v.Value), nil
		}
	}

	return "", nil
}

// SearchGeoIPInformationCase поиск объекта типа 'case' по его rootId
// возвращает _id объекта под которым он находится в БД и объект типа
func (dbs *DatabaseStorage) SearchGeoIPInformationCase(ctx context.Context, indexName, rootId string) (string, []IpAddressesInformation, error) {
	geoIpInformation := make([]IpAddressesInformation, 0)
	query := strings.NewReader(fmt.Sprintf("{\"query\": {\"bool\": {\"must\": [{\"match\": {\"event.rootId\": \"%s\"}}]}}}", rootId))

	//выполняем поиск _id индекса
	res, err := dbs.client.Search(
		dbs.client.Search.WithContext(ctx),
		dbs.client.Search.WithIndex(indexName),
		dbs.client.Search.WithBody(query),
	)
	if err != nil {
		return "", geoIpInformation, err
	}
	defer res.Body.Close()

	bodyRes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", geoIpInformation, err
	}

	result, err := supportingfunctions.GetElementsFromJSON(ctx, bodyRes)
	if err != nil {
		return "", geoIpInformation, err
	}

	resTmp := CaseDBResponse{}
	if err = json.Unmarshal(bodyRes, &resTmp); err != nil {
		return "", geoIpInformation, err
	}

	var underlineId string
	for k, v := range result {
		if k == "hits.hits._id" {
			underlineId = fmt.Sprint(v.Value)
		}
	}

	for _, v := range resTmp.Options.Hits[0].Source.AdditionalInformation.IpAddresses {
		geoIpInformation = append(geoIpInformation, IpAddressesInformation{
			Ip:          v.Ip,
			City:        v.City,
			Country:     v.Country,
			CountryCode: v.CountryCode,
		})
	}

	return underlineId, geoIpInformation, nil
}
