package documentgenerator

import (
	"fmt"
	"strings"

	caseobservables "github.com/av-belyakov/objectsthehiveformat/caseobservables"
	casettps "github.com/av-belyakov/objectsthehiveformat/casettps"
	objectsthehiveformat "github.com/av-belyakov/objectsthehiveformat/eventcase"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

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

// GetSensorAdditionalInformation дополнительная информация
func (c *VerifiedCase) GetSensorAdditionalInformation() *AdditionalInformation {
	return &c.AdditionalInformation
}

// SetSensorAdditionalInformation дополнительная информация
func (c *VerifiedCase) SetSensorAdditionalInformation(sai AdditionalInformation) {
	c.AdditionalInformation = sai
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

//********* listSensorId ************

// Get возвращает список идентификаторов сенсоров
func (e *listSensorId) Get() []string {
	return e.sensors
}

// AddElem добавляет только уникальные элементы
func (e *listSensorId) AddElem(sensorId string) {
	for _, v := range e.sensors {
		if v == sensorId {
			return
		}
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
	for _, v := range e.ip {
		if v == ip {
			return
		}
	}

	e.ip = append(e.ip, ip)
}
