package documentgenerator

import (
	"fmt"
	"strings"

	alerttemplate "github.com/av-belyakov/objectsthehiveformat/alert"
	alertdetails "github.com/av-belyakov/objectsthehiveformat/alertdetails"
	alertobjects "github.com/av-belyakov/objectsthehiveformat/alertobject"
	common "github.com/av-belyakov/objectsthehiveformat/common"
	eventalert "github.com/av-belyakov/objectsthehiveformat/eventalert"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/handlers"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
)

// AlertGenerator генерирует верифицированный объект типа 'alert'. Вернет первым элементом основной ID,
// вторым проверенный объект типа 'alert', третьим список полей которые не были обработаны
func AlertGenerator(chInput <-chan interfaces.CustomJsonDecoder) (string, *VerifiedAlert, map[string]string) {
	var (
		rootId string
		// список не обработанных полей
		listRawFields map[string]string = make(map[string]string)

		//Финальный объект
		verifiedAlert *VerifiedAlert = NewVerifiedAlert()

		event        *eventalert.TypeEventForAlert   = eventalert.NewTypeEventForAlert()
		eventObjects *alertobjects.EventAlertObject  = alertobjects.NewEventAlertObject()
		eventDetails *alertdetails.EventAlertDetails = alertdetails.NewEventAlertDetails()

		eventObjectCustomFields common.CustomFields = common.CustomFields{}
		alertObjectCustomFields common.CustomFields = common.CustomFields{}

		alert *alerttemplate.TypeAlert = alerttemplate.NewTypeAlert()
	)

	//******************* Вспомогательный объект для Artifacts **********************
	supportAlertArtifact := handlers.NewSupportiveAlertArtifacts()

	//******************* Основные обработчики для Event **********************
	// ------ EVENT ------
	listHandlerEvent := handlers.NewListHandlerEventAlert(event)

	// ------ EVENT OBJECT ------
	listHandlerEventObject := handlers.NewListHandlerEventAlertObject(eventObjects)

	// ------ EVENT OBJECT CUSTOMFIELDS ------
	listHandlerEventObjectCustomFields := handlers.NewListHandlerEventObjectCustomFields(eventObjectCustomFields)

	// ------ EVENT DETAILS ------
	listHandlerEventDetails := handlers.NewListHandlerEventAlertDetails(eventDetails)

	//******************* Основные обработчики для Alert **********************
	// ------ ALERT ------
	listHandlerAlert := handlers.NewListHandlerAlert(alert)

	// ------ ALERT CUSTOMFIELDS ------
	listHandlerAlertCustomFields := handlers.NewListHandlerAlertCustomFields(alertObjectCustomFields)

	// ------ ALERT ARTIFACTS ------
	listHandlerAlertArtifacts := handlers.NewListHandlerAlertArtifacts(supportAlertArtifact)

	for data := range chInput {
		var handlerIsExist bool
		verifiedAlert.SetID(data.GetUUID())

		if source, ok := searchEventSource(data.GetFieldBranch(), data.GetValue()); ok {
			verifiedAlert.SetSource(source)

			continue
		}

		if data.GetFieldBranch() == "event.rootId" {
			rootId = fmt.Sprint(data.GetValue())

			continue
		}

		//************ Обработчики для Event ************
		//event element
		if lf, ok := listHandlerEvent[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.GetValue())
			}

			continue
		}

		//event.object element
		if lf, ok := listHandlerEventObject[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.GetValue())
			}

			continue
		}

		//event.object.customFields element
		if lf, ok := listHandlerEventObjectCustomFields[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.GetValue())
			}

			continue
		}

		//event.details element
		if lf, ok := listHandlerEventDetails[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.GetValue())
			}

			continue
		}

		//************ Обработчики для Alert ************
		//alert element
		if lf, ok := listHandlerAlert[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.GetValue())
			}

			continue
		}

		//alert.customFields
		if lf, ok := listHandlerAlertCustomFields[data.GetFieldBranch()]; ok {
			handlerIsExist = true

			for _, f := range lf {
				f(data.GetValue())
			}

			continue
		}

		//alert.artifacts
		if strings.Contains(data.GetFieldBranch(), "alert.artifacts.") {
			handlerIsExist = true

			if lf, ok := listHandlerAlertArtifacts[data.GetFieldBranch()]; ok {
				for _, f := range lf {
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

	if eventObjects.GetDescription() != "" {
		//получаем id сенсора из event.object.description (если он там есть)
		if findSensorId, err := getSensorIdFromDescription(eventObjects.GetDescription()); err == nil {
			eventObjects.SetValueTags("sensor:id", findSensorId)
			eventDetails.SetValueTags("sensor:id", findSensorId)
		}
	}

	if eventDetails.GetDescription() != "" {
		//получаем id сенсора из event.details.description (если он там есть)
		if findSensorId, err := getSensorIdFromDescription(eventDetails.GetDescription()); err == nil {
			eventObjects.SetValueTags("sensor:id", findSensorId)
			eventDetails.SetValueTags("sensor:id", findSensorId)
		}
	}

	if alert.GetDescription() != "" {
		//получаем id сенсора из alert.description (если он там есть)
		if findSensorId, err := getSensorIdFromDescription(alert.GetDescription()); err == nil {
			alert.SetValueTags("sensor:id", findSensorId)
		}
	}

	//Собираем объект Alert
	eventObjects.SetValueCustomFields(eventObjectCustomFields)

	event.SetValueObject(*eventObjects)
	event.SetValueDetails(*eventDetails)

	alert.SetValueCustomFields(alertObjectCustomFields)
	alert.SetValueArtifacts(supportAlertArtifact.GetArtifacts())

	verifiedAlert.SetEvent(*event)
	verifiedAlert.SetAlert(*alert)

	//это если надо будеть обогащать alerts информацией по ip и сенсорам
	//objectElement := newDocument.Get().GetEvent().GetObject()
	//if listSensorId, ok := objectElement.GetTags()["sensor:id"]; ok {
	//	for _, v := range listSensorId {
	//		sensorsId.addElem(v)
	//	}
	//}

	//detailElem := newDocument.Get().GetEvent().GetDetails()
	//if listSensorId, ok := detailElem.GetTags()["sensor:id"]; ok {
	//	for _, v := range listSensorId {
	//		sensorsId.addElem(v)
	//	}
	//}

	return rootId, verifiedAlert, listRawFields
}
