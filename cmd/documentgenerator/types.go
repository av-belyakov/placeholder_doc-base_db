package documentgenerator

import (
	caseobservables "github.com/av-belyakov/objectsthehiveformat/caseobservables"
	casettps "github.com/av-belyakov/objectsthehiveformat/casettps"
	eventcase "github.com/av-belyakov/objectsthehiveformat/eventcase"

	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/countermessage"
)

type DocumentGenerator struct {
	logger  interfaces.Logger
	counter *countermessage.CounterMessage
	chInput <-chan interfaces.CustomJsonDecoder
}

// VerifiedCase объект представляет собой верифицированный 'case'
type VerifiedCase struct {
	ID              string                     `json:"@id" bson:"@id"`
	ElasticsearchID string                     `json:"@es_id" bson:"@es_id"`
	CreateTimestamp string                     `json:"@timestamp" bson:"@timestamp"`
	Source          string                     `json:"source" bson:"source"`
	Event           eventcase.TypeEventForCase `json:"event" bson:"event"`
	caseobservables.Observables
	casettps.Ttps
	AdditionalInformation
}

// AdditionalInformation дополнительная информация добавляемая к информации по кейсам
type AdditionalInformation struct {
	Sensors     []SensorInformation      `json:"@sensorAdditionalInformation"`
	IpAddresses []IpAddressesInformation `json:"@ipAddressAdditionalInformation"`
}

// SensorInformation содержит дополнительную информацию о сенсоре
type SensorInformation struct {
	SensorId    string `json:"sensorId" bson:"sensorId"`       //идентификатор сенсора
	HostId      string `json:"hostId" bson:"hostId"`           //идентификатор сенсора, специальный, для поиска информации в НКЦКИ
	GeoCode     string `json:"geoCode" bson:"geoCode"`         //географический код страны
	ObjectArea  string `json:"objectArea" bson:"objectArea"`   //сфера деятельности объекта
	SubjectRF   string `json:"subjectRF" bson:"subjectRF"`     //субъект Российской Федерации
	INN         string `json:"inn" bson:"inn"`                 //налоговый идентификатор
	HomeNet     string `json:"homeNet" bson:"homeNet"`         //перечень домашних сетей
	OrgName     string `json:"orgName" bson:"orgName"`         //наименование организации
	FullOrgName string `json:"fullOrgName" bson:"fullOrgName"` //полное наименование организации
}

// IpAddressesInformation дополнительная информация об ip адресе
type IpAddressesInformation struct {
	Ip          string `json:"ip"`          //ip адрес по которому выполнялся поиск
	City        string `json:"city"`        //город
	Country     string `json:"country"`     //страна
	CountryCode string `json:"countryCode"` //код страны
}
