package documentgenerator

import (
	casedetailshf "github.com/av-belyakov/objectsthehiveformat/casedetails"
	caseobjecthf "github.com/av-belyakov/objectsthehiveformat/caseobject"
	commonhf "github.com/av-belyakov/objectsthehiveformat/common"
	eventcasehf "github.com/av-belyakov/objectsthehiveformat/eventcase"

	"github.com/av-belyakov/placeholder_doc-base_db/cmd/handlers"
	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
)

// CaseGenerator генерирует верифицированный объект типа 'case'
func (dg *DocumentGenerator) CaseGenerator(chInput <-chan interfaces.CustomJsonDecoder) VerifiedCase {
	var (
		rootId string
		// список не обработанных полей
		listRawFields map[string]string = make(map[string]string)

		//Финальный объект
		verifiedCase VerifiedCase = VerifiedCase{}

		event        *eventcasehf.TypeEventForCase   = eventcasehf.NewTypeEventForCase()
		eventObjects *caseobjecthf.EventCaseObject   = caseobjecthf.NewEventCaseObject()
		eventDetails *casedetailshf.EventCaseDetails = casedetailshf.NewEventCaseDetails()

		eventObjectCustomFields  commonhf.CustomFields = commonhf.CustomFields{}
		eventDetailsCustomFields commonhf.CustomFields = commonhf.CustomFields{}
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

	return verifiedCase
}
