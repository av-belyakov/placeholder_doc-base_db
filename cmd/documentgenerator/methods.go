package documentgenerator

import (
	"fmt"
	"strings"
	"time"

	alert "github.com/av-belyakov/objectsthehiveformat/alert"
	caseobservables "github.com/av-belyakov/objectsthehiveformat/caseobservables"
	casettps "github.com/av-belyakov/objectsthehiveformat/casettps"
	eventalert "github.com/av-belyakov/objectsthehiveformat/eventalert"
	objectsthehiveformat "github.com/av-belyakov/objectsthehiveformat/eventcase"

	"slices"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

//********** VerifiedAlert ************

// NewVerifiedAlert новый элемент содержащий проверенный объект типа 'case'
func NewVerifiedAlert() *VerifiedAlert {
	return &VerifiedAlert{
		Alert:           *alert.NewTypeAlert(),
		Event:           *eventalert.NewTypeEventForAlert(),
		CreateTimestamp: supportingfunctions.GetDateTimeFormatRFC3339(time.Now().UnixMilli()),
	}
}

func (a *VerifiedAlert) Get() *VerifiedAlert {
	return a
}

// GetID уникальный идентификатор
func (va *VerifiedAlert) GetID() string {
	return va.ID
}

// SetID уникальный идентификатор
func (va *VerifiedAlert) SetID(id string) {
	va.ID = id
}

// GetCreateTimestamp время создания объекта в формате RFC3339
func (va *VerifiedAlert) GetCreateTimestamp() string {
	return va.CreateTimestamp
}

// SetCreateTimestamp время создания объекта в формате RFC3339
func (va *VerifiedAlert) SetCreateTimestamp(time string) {
	va.CreateTimestamp = time
}

// GetSource наименование источника
func (va *VerifiedAlert) GetSource() string {
	return va.Source
}

// SetSource наименование источника
func (va *VerifiedAlert) SetSource(source string) {
	va.Source = source
}

// GetEvent объект типа 'event'
func (va *VerifiedAlert) GetEvent() *eventalert.TypeEventForAlert {
	return &va.Event
}

// SetEvent объект типа 'event'
func (va *VerifiedAlert) SetEvent(event eventalert.TypeEventForAlert) {
	va.Event = event
}

// GetAlert объект типа 'alert'
func (va *VerifiedAlert) GetAlert() *alert.TypeAlert {
	return &va.Alert
}

// SetAlert объект типа 'alert'
func (va *VerifiedAlert) SetAlert(alert alert.TypeAlert) {
	va.Alert = alert
}

func (va *VerifiedAlert) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	strB := strings.Builder{}
	strB.WriteString(fmt.Sprintf("%s'createTimestatmp': '%s'\n", ws, va.CreateTimestamp))
	strB.WriteString(fmt.Sprintf("%s'source': '%s'\n", ws, va.Source))
	strB.WriteString(fmt.Sprintf("%s'event':\n", ws))
	strB.WriteString(va.Event.ToStringBeautiful(num + 1))
	strB.WriteString(fmt.Sprintf("%s'alert':\n", ws))
	strB.WriteString(va.Alert.ToStringBeautiful(num + 1))

	return strB.String()
}

//********** VerifiedCase ***********

// NewVerifiedCase новый элемент содержащий проверенный объект типа 'case'
func NewVerifiedCase() *VerifiedCase {
	return &VerifiedCase{}
}

func (c *VerifiedCase) Get() *VerifiedCase {
	return c
}

// GetID идентификатор объекта
func (vc *VerifiedCase) GetID() string {
	return vc.ID
}

// SetID идентификатор объекта
func (vc *VerifiedCase) SetID(v string) {
	vc.ID = v
}

// GetSource наименование источника
func (vc *VerifiedCase) GetSource() string {
	return vc.Source
}

// SetSource наименование источника
func (vc *VerifiedCase) SetSource(v string) {
	vc.Source = v
}

// GetCreateTimestamp временная метка
func (c *VerifiedCase) GetCreateTimestamp() string {
	return c.CreateTimestamp
}

// SetCreateTimestamp временная метка
func (c *VerifiedCase) SetCreateTimestamp(timestamp string) {
	c.CreateTimestamp = timestamp
}

// GetEvent объект типа 'event'
func (c *VerifiedCase) GetEvent() *objectsthehiveformat.TypeEventForCase {
	return &c.Event
}

// SetEvent объект типа 'event'
func (c *VerifiedCase) SetEvent(v objectsthehiveformat.TypeEventForCase) {
	c.Event = v
}

// GetObservables объект типа 'observables'
func (c *VerifiedCase) GetObservables() *caseobservables.Observables {
	return &c.Observables
}

// SetObservables объект типа 'observables'
func (c *VerifiedCase) SetObservables(v caseobservables.Observables) {
	c.Observables = v
}

// GetTtps объект типа 'ttps'
func (c *VerifiedCase) GetTtps() *casettps.Ttps {
	return &c.Ttps
}

// SetTtps объект типа 'ttps'
func (c *VerifiedCase) SetTtps(v casettps.Ttps) {
	c.Ttps = v
}

// GetAdditionalInformation дополнительная информация
func (c *VerifiedCase) GetAdditionalInformation() *AdditionalInformation {
	return &c.AdditionalInformation
}

// SetAdditionalInformation дополнительная информация
func (c *VerifiedCase) SetAdditionalInformation(sai AdditionalInformation) {
	c.AdditionalInformation = sai
}

// GetSensorsInformation объекты с информацией о сенсоре
func (ai *AdditionalInformation) GetSensorsInformation() []SensorInformation {
	return ai.Sensors
}

// AddSensorInformation добавляет информацию о сенсоре
func (ai *AdditionalInformation) AddSensorInformation(v SensorInformation) {
	if slices.ContainsFunc(ai.Sensors, func(obj SensorInformation) bool {
		return obj.SensorId == v.SensorId
	}) {
		return
	}

	ai.Sensors = append(ai.Sensors, v)
}

// GetIpAddressesInformation объекты с информацией об ip адресах
func (ai *AdditionalInformation) GetIpAddressesInformation() []IpAddressInformation {
	return ai.IpAddresses
}

// AddGetIpAddressInformation добавляет информацию об ip адресе
func (ai *AdditionalInformation) AddGetIpAddressInformation(v IpAddressInformation) {
	if slices.ContainsFunc(ai.IpAddresses, func(obj IpAddressInformation) bool {
		return obj.Ip == v.Ip
	}) {
		return
	}

	ai.IpAddresses = append(ai.IpAddresses, v)
}

// GetIpAddrString ip адрес в виде строки
func (i IpAddressInformation) GetIpAddrString() string {
	return i.Ip
}

func (c *VerifiedCase) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("%s'@id': '%s'\n", ws, c.ID))
	str.WriteString(fmt.Sprintf("%s'@createTimestatmp': '%s'\n", ws, c.CreateTimestamp))
	str.WriteString(fmt.Sprintf("%s'source': '%s'\n", ws, c.Source))
	str.WriteString(fmt.Sprintf("%s'event':\n", ws))
	str.WriteString(c.Event.ToStringBeautiful(num + 1))
	str.WriteString(c.Observables.ToStringBeautiful(num))
	str.WriteString(c.Ttps.ToStringBeautiful(num))

	return str.String()
}

// ToStringBeautiful дополнительная информация по сенсорам и ip адресам
func (ai *AdditionalInformation) ToStringBeautiful(num int) string {
	var str strings.Builder = strings.Builder{}
	str.WriteString(fmt.Sprintf("%s'@sensorAdditionalInformation':\n", supportingfunctions.GetWhitespace(num)))
	for k, v := range ai.Sensors {
		str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+1), k+1))
		str.WriteString(v.ToStringBeautiful(num + 2))
	}
	str.WriteString(fmt.Sprintf("%s'@ipAddressAdditionalInformation':\n", supportingfunctions.GetWhitespace(num)))
	for k, v := range ai.IpAddresses {
		str.WriteString(fmt.Sprintf("%s%d.\n", supportingfunctions.GetWhitespace(num+1), k+1))
		str.WriteString(v.ToStringBeautiful(num + 2))
	}

	return str.String()
}

// GetSensorId идентификатор сенсора
func (si *SensorInformation) GetSensorId() string {
	return si.SensorId
}

// ToStringBeautiful для информации по сенсору
func (si *SensorInformation) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("%s'sensorId': '%s'\n", ws, si.SensorId))
	str.WriteString(fmt.Sprintf("%s'hostId': '%s'\n", ws, si.HostId))
	str.WriteString(fmt.Sprintf("%s'geoCode': '%s'\n", ws, si.GeoCode))
	str.WriteString(fmt.Sprintf("%s'objectArea': '%s'\n", ws, si.ObjectArea))
	str.WriteString(fmt.Sprintf("%s'subjectRF': '%s'\n", ws, si.SubjectRF))
	str.WriteString(fmt.Sprintf("%s'inn': '%s'\n", ws, si.INN))
	str.WriteString(fmt.Sprintf("%s'homeNet': '%s'\n", ws, si.HomeNet))
	str.WriteString(fmt.Sprintf("%s'orgName': '%s'\n", ws, si.OrgName))
	str.WriteString(fmt.Sprintf("%s'fullOrgName': '%s'\n", ws, si.FullOrgName))

	return str.String()
}

// ToStringBeautiful для информации по ip адресу
func (i *IpAddressInformation) ToStringBeautiful(num int) string {
	ws := supportingfunctions.GetWhitespace(num)

	str := strings.Builder{}
	str.WriteString(fmt.Sprintf("%s'Ip': '%s'\n", ws, i.Ip))
	str.WriteString(fmt.Sprintf("%s'City': '%s'\n", ws, i.City))
	str.WriteString(fmt.Sprintf("%s'Country': '%s'\n", ws, i.Country))
	str.WriteString(fmt.Sprintf("%s'CountryCode': '%s'\n", ws, i.CountryCode))

	return str.String()
}

//********* listSensorId ************

// Get возвращает список идентификаторов сенсоров
func (e *listSensorId) Get() []string {
	return e.sensors
}

// AddElem добавляет только уникальные элементы
func (e *listSensorId) AddElem(sensorId string) {
	if slices.Contains(e.sensors, sensorId) {
		return
	}

	e.sensors = append(e.sensors, sensorId)
}

//********* listIpAddresses ************

// Get возвращает список ip
func (e *listIpAddresses) Get() []string {
	return e.ip
}

// AddElem добавляет только уникальные элементы
func (e *listIpAddresses) AddElem(ip string) {
	if slices.Contains(e.ip, ip) {
		return
	}

	e.ip = append(e.ip, ip)
}
