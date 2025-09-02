package documentgenerator

import (
	alert "github.com/av-belyakov/objectsthehiveformat/alert"
	caseobservables "github.com/av-belyakov/objectsthehiveformat/caseobservables"
	casettps "github.com/av-belyakov/objectsthehiveformat/casettps"
	eventalert "github.com/av-belyakov/objectsthehiveformat/eventalert"
	eventcase "github.com/av-belyakov/objectsthehiveformat/eventcase"
)

// VerifiedAlert объект представляет собой верифицированный тип 'alert'
type VerifiedAlert struct {
	Event           eventalert.TypeEventForAlert `json:"event" bson:"event"`
	Alert           alert.TypeAlert              `json:"alert,omitzero" bson:"alert"`
	ID              string                       `json:"@id" bson:"@id"`
	Source          string                       `json:"source" bson:"source"`
	CreateTimestamp string                       `json:"@timestamp" bson:"@timestamp"`
	ElasticsearchID string                       `json:"@es_id" bson:"@es_id"`
}

// VerifiedCase объект представляет собой верифицированный тип 'case'
type VerifiedCase struct {
	Event eventcase.TypeEventForCase `json:"event" bson:"event"`
	caseobservables.Observables
	casettps.Ttps
	AdditionalInformation
	ID              string `json:"@id" bson:"@id"`
	Source          string `json:"source" bson:"source"`
	ElasticsearchID string `json:"@es_id" bson:"@es_id"`
	CreateTimestamp string `json:"@timestamp" bson:"@timestamp"`
}

// AdditionalInformation дополнительная информация добавляемая к информации по кейсам
type AdditionalInformation struct {
	Sensors     []SensorInformation    `json:"@sensorAdditionalInformation"`
	IpAddresses []IpAddressInformation `json:"@ipAddressAdditionalInformation"`
}

// SensorInformation содержит дополнительную информацию о сенсоре
type SensorInformation struct {
	INN         string `json:"inn" bson:"inn"`                 //налоговый идентификатор
	HostId      string `json:"hostId" bson:"hostId"`           //идентификатор сенсора, специальный, для поиска информации в НКЦКИ
	OrgName     string `json:"orgName" bson:"orgName"`         //наименование организации
	HomeNet     string `json:"homeNet" bson:"homeNet"`         //перечень домашних сетей
	GeoCode     string `json:"geoCode" bson:"geoCode"`         //географический код страны
	SensorId    string `json:"sensorId" bson:"sensorId"`       //идентификатор сенсора
	SubjectRF   string `json:"subjectRF" bson:"subjectRF"`     //субъект Российской Федерации
	ObjectArea  string `json:"objectArea" bson:"objectArea"`   //сфера деятельности объекта
	FullOrgName string `json:"fullOrgName" bson:"fullOrgName"` //полное наименование организации
}

// IpAddressesInformation дополнительная информация об ip адресе
type IpAddressInformation struct {
	Ip          string `json:"ip"`          //ip адрес по которому выполнялся поиск
	City        string `json:"city"`        //город
	Country     string `json:"country"`     //страна
	CountryCode string `json:"countryCode"` //код страны
}

type listSensorId struct {
	sensors []string
}

type listIpAddresses struct {
	ip []string
}
