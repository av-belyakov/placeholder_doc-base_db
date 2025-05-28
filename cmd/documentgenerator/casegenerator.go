package documentgenerator

import (
	"fmt"
	"reflect"

	casedetails "github.com/av-belyakov/objectsthehiveformat/casedetails"
	caseobjects "github.com/av-belyakov/objectsthehiveformat/caseobject"
	caseobservables "github.com/av-belyakov/objectsthehiveformat/caseobservables"
	casettps "github.com/av-belyakov/objectsthehiveformat/casettps"
	common "github.com/av-belyakov/objectsthehiveformat/common"
	eventcase "github.com/av-belyakov/objectsthehiveformat/eventcase"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/handlers"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
)

// CaseGenerator генерирует верифицированный объект типа 'case'. Вернет первым элементом основной ID,
// вторым проверенный объект типа 'case', третьим список полей которые не были обработаны
func CaseGenerator(chInput <-chan interfaces.CustomJsonDecoder) (string, *VerifiedCase, map[string]string) {
	var (
		rootId string

		// список не обработанных полей
		listRawFields map[string]string = make(map[string]string)

		//объект с дополнительной информацией по сенсорам и ip адресам
		additionalInformation AdditionalInformation = AdditionalInformation{
			Sensors:     []SensorInformation{},
			IpAddresses: []IpAddressInformation{},
		}

		//Финальный объект
		verifiedCase *VerifiedCase = NewVerifiedCase()

		event        *eventcase.TypeEventForCase   = eventcase.NewTypeEventForCase()
		eventObjects *caseobjects.EventCaseObject  = caseobjects.NewEventCaseObject()
		eventDetails *casedetails.EventCaseDetails = casedetails.NewEventCaseDetails()

		eventObjectCustomFields  common.CustomFields = common.CustomFields{}
		eventDetailsCustomFields common.CustomFields = common.CustomFields{}
	)

	//******************* Основные обработчики для Event **********************
	// ------ EVENT ------
	listHandlerEvent := handlers.NewListHandlerEventCase(event)
	// ------ EVENT OBJECT ------
	listHandlerEventObject := handlers.NewListHandlerEventCaseObject(eventObjects)
	// ------ EVENT OBJECT CUSTOMFIELDS ------
	listHandlerEventObjectCustomFields := handlers.NewListHandlerEventObjectCustomFields(eventObjectCustomFields)
	// ------ EVENT DETAILS ------
	listHandlerEventDetails := handlers.NewListHandlerEventCaseDetails(eventDetails)
	// ------ EVENT DETAILS CUSTOMFIELDS ------
	listHandlerEventDetailsCustomFields := handlers.NewListHandlerEventObjectCustomFields(eventDetailsCustomFields)

	//******************* Вспомогательный объект для Observables **********************
	so := handlers.NewSupportiveObservables()
	listHandlerObservables := handlers.NewListHandlerObservables(so)

	//******************* Вспомогательный объект для Ttp **********************
	sttp := handlers.NewSupportiveTtp()
	listHandlerTtp := handlers.NewListHandlerTtp(sttp)

	for data := range chInput {
		var handlerIsExist bool
		verifiedCase.SetID(data.GetUUID())

		if source, ok := searchEventSource(data.GetFieldBranch(), data.GetValue()); ok {
			verifiedCase.SetSource(source)

			continue
		}

		if data.GetFieldBranch() == "event.rootId" {
			rootId = fmt.Sprint(data.GetValue())

			continue
		}

		//******************************************************************
		//********** Сбор всех объектов относящихся к полю Event  **********
		// event element
		if lf, ok := listHandlerEvent[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.GetValue())
			}

			continue
		}

		// event.object element
		if lf, ok := listHandlerEventObject[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.GetValue())
			}

			continue
		}

		// event.object.customFields element
		if lf, ok := listHandlerEventObjectCustomFields[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.GetValue())
			}

			continue
		}

		// event.details element
		if lf, ok := listHandlerEventDetails[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.GetValue())
			}

			continue
		}

		// event.details.customFields element
		if lf, ok := listHandlerEventDetailsCustomFields[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.GetValue())
			}

			continue
		}

		//************************************************************************
		//********** Сбор всех объектов относящихся к полю Observables  **********
		// для всех полей входящих в observables, кроме содержимого поля reports
		if lf, ok := listHandlerObservables[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				r := reflect.TypeOf(data.GetValue())
				switch r.Kind() {
				case reflect.Slice:
					if s, ok := data.GetValue().([]any); ok {
						for _, value := range s {
							f(value)
						}
					}

				default:
					f(data.GetValue())

				}
			}

			continue
		}

		//убрал обработку observables.reports так как тип TtpsMessageEs
		//способствует росту черезмерно большого количества полей которое
		//влечет за собой превышения лимита маппинга в Elsticsearch), что
		//выражается в ошибке от СУБД типа "Limit of total fields [2000]
		//has been exceeded while adding new fields"
		//
		//для всех полей входящих в состав observables.reports
		//if strings.Contains(data.FieldBranch, "observables.reports.") {
		//		handlerIsExist = true
		//		so.HandlerReportValue(data.FieldBranch, data.Value)
		//}

		//*********************************************************************
		//********** Сбор всех объектов относящихся к полю Ttp  ***************
		if lf, ok := listHandlerTtp[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				r := reflect.TypeOf(data.GetValue())
				switch r.Kind() {
				case reflect.Slice:
					if s, ok := data.GetValue().([]any); ok {
						for _, value := range s {
							f(value)
						}
					}

				default:
					f(data.GetValue())

				}
			}

			continue
		}

		if !handlerIsExist {
			// записываем в лог-файл поля, которые не были обработаны
			listRawFields[data.GetFieldBranch()] = fmt.Sprint(data.GetValue())
		}
	}

	//проверяем что, в поле event.object.customFields.first-time временное значение
	// не соответствует 1970-01-01T00:00:00+00:00, так как это равносильно пустому
	//значению если значение поля больше чем 1970-01-01T00:00:00+00:00, то в
	//@timestamp укладываем его, иначе используем значение из поля event.object._createAt
	verifiedCase.SetCreateTimestamp(eventObjects.GetCreatedAt())
	for k, v := range eventObjectCustomFields {
		if k == "first-time" {
			_, _, _, date := v.Get()
			if date != "1970-01-01T00:00:00+00:00" {
				verifiedCase.SetCreateTimestamp(date)
			}

			break
		}
	}

	//собираем объект Event
	eventObjects.SetValueCustomFields(eventObjectCustomFields)
	eventDetails.SetValueCustomFields(eventDetailsCustomFields)
	event.SetValueObject(*eventObjects)
	event.SetValueDetails(*eventDetails)

	// собираем объект observables
	observables := caseobservables.NewObservables()
	observables.SetValueObservables(so.GetObservables())

	// собираем объект ttp
	ttps := casettps.NewTtps()
	ttps.SetTtps(sttp.GetTtps())

	verifiedCase.SetEvent(*event)
	verifiedCase.SetObservables(*observables)
	verifiedCase.SetTtps(*ttps)

	eventCase := verifiedCase.GetEvent()
	objectElem := eventCase.GetObject()

	//формируется список идентификаторов сенсоров
	if listSensorId, ok := objectElem.GetTags()["sensor:id"]; ok {
		for _, v := range listSensorId {
			additionalInformation.AddSensorInformation(SensorInformation{
				SensorId: v,
			})
		}
	}

	//формируется список ip адресов
	if ipObservables, ok := verifiedCase.GetKeyObservables("ip"); ok {
		for _, v := range ipObservables {
			additionalInformation.AddGetIpAddressInformation(IpAddressInformation{
				Ip: v.Data,
			})
		}
	}
	if listIpAddresses, ok := objectElem.GetTags()["ip"]; ok {
		for _, v := range listIpAddresses {
			additionalInformation.AddGetIpAddressInformation(IpAddressInformation{
				Ip: v,
			})
		}
	}

	verifiedCase.SetAdditionalInformation(additionalInformation)

	return rootId, verifiedCase, listRawFields
}
