package databasestorageapi

import (
	"encoding/json"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/request"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

// sendGeoIpRequest запрос для обогащения кейса дополнительной информацией о
// локальном местоположении ip адресов
func sendGeoIpRequest(
	id string,
	list []string,
	getChan func() chan SettingsChanOutput) (bool, error) {
	if len(list) == 0 {
		return false, nil
	}

	reqGeoIP, err := json.Marshal(request.RequestGeoIP{
		Source:          "placeholder_doc-base_db",
		TaskId:          id,
		ListIpAddresses: list,
	})
	if err != nil {
		return false, err
	}

	getChan() <- SettingsChanOutput{
		Command: "get_geoip_info",
		RootId:  id,
		Data:    reqGeoIP,
	}

	return true, nil
}

// sendSensorInformationRequest запрос для обогащения кейса дополнительной информацией по
// расположению сенсоров
func sendSensorInformationRequest(
	id string,
	list []string,
	getChan func() chan SettingsChanOutput) (bool, error) {
	if len(list) == 0 {
		return false, nil
	}

	reqSensorId, err := json.Marshal(request.RequestSensorInformation{
		Source:      "placeholder_doc-base_db",
		TaskId:      id,
		ListSensors: list,
	})
	if err != nil {
		return false, err
	}

	getChan() <- SettingsChanOutput{
		Command: "get_sensor_info",
		RootId:  id,
		Data:    reqSensorId,
	}

	return true, nil
}

// bodyClose закрывает ответ с предварительной проверкой
func bodyClose(res *esapi.Response) {
	if res == nil || res.Body == nil {
		return
	}

	res.Body.Close()
}
