package databasestorageapi

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/response"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// addGeoIPInformation дополнительная информация о географическом местоположении ip адресов
func (dbs *DatabaseStorage) addGeoIPInformation(ctx context.Context, data any) {
	newData, ok := data.([]byte)
	if !ok {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("type conversion error")).Error())

		return
	}

	var newDocument response.ResponseGeoIpInformation
	if err := json.Unmarshal(newData, &newDocument); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	dbs.logger.Send("info", fmt.Sprintf("section:'information handling', command:'add geoip information', accepted object:'%#v'", newDocument))

	//если в принятом ответе от модуля обогащения информацией о географическом местоположении
	//ip адресов есть глобальная ошибка
	if newDocument.Error != "" {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(newDocument.Error)).Error())

		return
	}

	//получаем наименование хранилища
	indexName, isExist := dbs.settings.storages["case"]
	if !isExist {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	t := time.Now()
	month := int(t.Month())
	//текущий индекс
	indexCurrent := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), month)

	//добавляем небольшую задержку что бы СУБД успела добавить индекс
	//***************************************************************
	time.Sleep(3 * time.Second)
	//***************************************************************

	//поиск _id объекта типа 'case' по его rootId (что в передается в newDocument.TaskId)
	underlineId, err := dbs.SearchUnderlineIdCase(ctx, indexCurrent, newDocument.TaskId)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	//формируется список с информацией по ip адресам
	var ipInfoList []IpAddressesInformation
	for _, ipAddress := range newDocument.Informations {
		if ipAddress.Error != "" {
			dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(ipAddress.Error)).Error())

			ipInfoList = append(ipInfoList, IpAddressesInformation{
				Ip: ipAddress.IpAddr,
			})

			continue
		}

		ipInfoList = append(ipInfoList, IpAddressesInformation{
			Ip:          ipAddress.IpAddr,
			City:        ipAddress.City,
			Country:     ipAddress.Country,
			CountryCode: ipAddress.Code,
		})
	}

	//обновляемый список не должен быть пустым, а то получается что пустой
	//список с ip адресами затирается вообще пустым списком
	if len(ipInfoList) == 0 {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the list with information on ip addresses should not be empty")).Error())

		return
	}

	request, err := json.MarshalIndent(AdditionalInformationIpAddress{
		IpAddresses: ipInfoList,
	}, "", " ")
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%w'", newDocument.TaskId, err)).Error())

		return
	}

	//обновление информации в БД
	bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))
	res, err := dbs.client.Update(indexCurrent, underlineId, bodyUpdate)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%w'", newDocument.TaskId, err)).Error())

		return
	}
	defer res.Body.Close()

	if res != nil && res.StatusCode != http.StatusOK {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%w'", newDocument.TaskId, err)).Error())

		return
	}

	dbs.logger.Send("info", fmt.Sprintf("the geoip information for root id:'%s' it was added successfully", newDocument.TaskId))
}

// addSensorInformation дополнительная информация о местоположении и принадлежности сенсоров
func (dbs *DatabaseStorage) addSensorInformation(ctx context.Context, data any) {
	newData, ok := data.([]byte)
	if !ok {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("type conversion error")).Error())

		return
	}

	var newDocument response.ResponseSensorsInformation
	if err := json.Unmarshal(newData, &newDocument); err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(err).Error())

		return
	}

	dbs.logger.Send("info", fmt.Sprintf("section:'information handling', command:'add sensor information', accepted object:'%#v'", newDocument))

	//если в принятом ответе от модуля обогащения информацией о сенсорах есть глобальная ошибка
	if newDocument.Error != "" {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(newDocument.Error)).Error())

		return
	}

	//получаем наименование хранилища
	indexName, isExist := dbs.settings.storages["case"]
	if !isExist {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	t := time.Now()
	month := int(t.Month())
	//текущий индекс
	indexCurrent := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), month)

	//добавляем небольшую задержку что бы СУБД успела добавить индекс
	//***************************************************************
	time.Sleep(3 * time.Second)
	//***************************************************************

	//поиск _id объекта типа 'case' по его rootId (что в передается в newDocument.TaskId)
	underlineId, err := dbs.SearchUnderlineIdCase(ctx, indexCurrent, newDocument.TaskId)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	//формируется список с информацией по сенсорам
	var sensorInfoList []SensorInformation
	for _, sensor := range newDocument.Informations {
		if sensor.Error != "" {
			dbs.logger.Send("error", supportingfunctions.CustomError(errors.New(sensor.Error)).Error())
		}

		sensorInfoList = append(sensorInfoList, SensorInformation{
			SensorId:    sensor.SensorID,
			GeoCode:     sensor.GeoCode,
			ObjectArea:  sensor.ObjectArea,
			SubjectRF:   sensor.SubjectRussianFederation,
			INN:         sensor.INN,
			HomeNet:     sensor.HomeNet,
			OrgName:     sensor.OrganizationName,
			FullOrgName: sensor.FullOrganizationName,
		})
	}

	//обновляемый список не должен быть пустым, а то получается что пустой
	//список с информацией по сенсорам затирается вообще пустым списком
	if len(sensorInfoList) == 0 {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the list with information on sensors should not be empty")).Error())

		return
	}

	request, err := json.MarshalIndent(AdditionalInformationSensors{Sensors: sensorInfoList}, "", " ")
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%w'", newDocument.TaskId, err)).Error())

		return
	}

	//обновление информации в БД
	bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))
	res, err := dbs.client.Update(indexCurrent, underlineId, bodyUpdate)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%w'", newDocument.TaskId, err)).Error())

		return
	}
	defer res.Body.Close()

	if res != nil && res.StatusCode != http.StatusOK {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%w'", newDocument.TaskId, err)).Error())

		return
	}

	dbs.logger.Send("info", fmt.Sprintf("the sensors information for root id:'%s' it was added successfully", newDocument.TaskId))
}
