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
	//добавляем небольшую задержку что бы СУБД успела добавить индекс
	//***************************************************************
	time.Sleep(3 * time.Second)
	//***************************************************************

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

	//поиск _id объекта типа 'case' по его rootId (что в передается в newDocument.TaskId)
	underlineId, err := dbs.SearchUnderlineIdCase(ctx, indexCurrent, newDocument.TaskId)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(errors.New("the identifier of the index name was not found")).Error())

		return
	}

	//формируется список с информацией по ip адресам
	var ipInfoList []IpAddressesInformation
	for _, ipAddress := range newDocument.Informations {
		ipInfoList = append(ipInfoList, IpAddressesInformation{
			Ip:          ipAddress.IpAddr,
			City:        ipAddress.City,
			Country:     ipAddress.Country,
			CountryCode: ipAddress.Code,
		})
	}

	request, err := json.MarshalIndent(AdditionalInformationIpAddress{
		IpAddresses: ipInfoList,
	}, "", " ")
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%s'", newDocument.TaskId, err.Error())).Error())

		return
	}

	//обновление информации в БД
	bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))
	res, err := dbs.client.Update(indexCurrent, underlineId, bodyUpdate)
	if err != nil {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%s'", newDocument.TaskId, err.Error())).Error())

		return
	}
	defer res.Body.Close()

	if res != nil && res.StatusCode != http.StatusOK {
		dbs.logger.Send("error", supportingfunctions.CustomError(fmt.Errorf("'rootId:'%s', '%s'", newDocument.TaskId, err.Error())).Error())
	}
}

// addSensorInformation дополнительная информация о местоположении и принадлежности сенсоров
func (dbs *DatabaseStorage) addSensorInformation(ctx context.Context, data any) {
	/*
	   //добавляем небольшую задержку что бы СУБД успела добавить индекс
	   //***************************************************************
	   time.Sleep(3 * time.Second)
	   //***************************************************************

	   newDocument, ok := data.(response.ResponseSensorsInformation)

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

	   // SearchUnderlineIdCase поиск объекта типа 'case' по его _id
	   underlineId, err := dbs.SearchUnderlineIdCase(ctx, indexName, newDocument.TaskId)

	   /*

	   	Здесь надо как то найти кейс с соответствующим ip и обновить у него информацию
	   	об GeoIP ip адресах
	*/
}

// AddEventenrichmentCase выполняет обогащение уже имеющегося кейса дополнительной, полезной информацией
/*func (hsd HandlerSendData) AddEventenrichmentCase(data interface{}, indexName string, logging chan<- datamodels.MessageLogging) {

	//добавляем небольшую задержку что бы СУБД успела добавить индекс
	//***************************************************************
	time.Sleep(3 * time.Second)
	//***************************************************************

	additionalInformation := datamodels.AdditionalInformation{}

	//приводим значение к интерфейсу позволяющему получить доступ к информации о сенсорах
	infoEvent, ok := data.(datamodels.InformationFromEventEnricher)
	if !ok {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'error converting the type to type datamodels.InformationFromEventEnricher' %s:%d", f, l-1),
			MsgType: "error",
		}

		return
	}

	t := time.Now()
	month := int(t.Month())
	indexCurrent := fmt.Sprintf("%s_%d_%d", indexName, t.Year(), month)

	//выполняем поиск _id индекса
	caseId, err := SearchUnderlineIdCase(indexCurrent, infoEvent.GetRootId(), infoEvent.GetSource(), hsd)
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'rootId: '%s', %s' %s:%d", infoEvent.GetRootId(), err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	//информация по сенсорам
	invalidString := "DOCTYPE"
	for _, v := range infoEvent.GetSensorsId() {
		//убираем невалидные данные
		if strings.Contains(infoEvent.GetGeoCode(v), invalidString) || strings.Contains(infoEvent.GetObjectArea(v), invalidString) || strings.Contains(infoEvent.GetSubjectRF(v), invalidString) || strings.Contains(infoEvent.GetINN(v), invalidString) {
			_, f, l, _ := runtime.Caller(0)
			logging <- datamodels.MessageLogging{
				MsgData: fmt.Sprintf("'sensor '%s' information contains incorrect data %s:%d", v, f, l-2),
				MsgType: "error",
			}

			continue
		}

		si := datamodels.NewSensorInformation()
		si.SetSensorId(v)
		si.SetHostId(infoEvent.GetHostId(v))
		si.SetGeoCode(infoEvent.GetGeoCode(v))
		si.SetObjectArea(infoEvent.GetObjectArea(v))
		si.SetSubjectRF(infoEvent.GetSubjectRF(v))
		si.SetINN(infoEvent.GetINN(v))
		si.SetHomeNet(infoEvent.GetHomeNet(v))
		si.SetOrgName(infoEvent.GetOrgName(v))
		si.SetFullOrgName(infoEvent.GetFullOrgName(v))

		additionalInformation.AddSensor(*si)
	}

	//информация по ip адресам
	for _, ipAddress := range infoEvent.GetIpAddresses() {
		if !infoEvent.GetIsSuccess(ipAddress) {
			continue
		}

		ipi := datamodels.NewIpAddressesInformation()
		ipi.SetIp(ipAddress)

		customIpInfo := groupIpInfoResult(infoEvent)

		ipi.SetCity(customIpInfo.city)
		ipi.SetCountry(customIpInfo.country)
		ipi.SetCountryCode(customIpInfo.countryCode)

		additionalInformation.AddIpAddress(*ipi)
	}

	request, err := json.MarshalIndent(*additionalInformation.Get(), "", " ")
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'rootId: '%s', '%s' %s:%d", infoEvent.GetRootId(), err.Error(), f, l-2),
			MsgType: "error",
		}

		return
	}

	bodyUpdate := strings.NewReader(fmt.Sprintf("{\"doc\": %s}", string(request)))
	res, err := hsd.Client.Update(indexCurrent, caseId, bodyUpdate)
	defer func() {
		if res == nil || res.Body == nil {
			return
		}

		errClose := res.Body.Close()
		if err == nil {
			err = errClose
		}
	}()
	if err != nil {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'rootId: '%s', %s' %s:%d", err.Error(), infoEvent.GetRootId(), f, l-1),
			MsgType: "error",
		}

		return
	}

	if res != nil && res.StatusCode != http.StatusOK {
		_, f, l, _ := runtime.Caller(0)
		logging <- datamodels.MessageLogging{
			MsgData: fmt.Sprintf("'rootId: '%s', %d %s' %s:%d", infoEvent.GetRootId(), res.StatusCode, res.Status(), f, l-1),
			MsgType: "error",
		}
	}
}*/
