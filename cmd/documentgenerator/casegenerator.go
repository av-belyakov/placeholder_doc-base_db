package documentgenerator

import (
	eventcase "github.com/av-belyakov/objectsthehiveformat/eventcase"
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

		event *eventcase.TypeEventForCase = eventcase.NewTypeEventForCase()
		// event *datamodels.EventMessageForEsCase = datamodels.NewEventMessageForEsCase()

		eventObject *objectsthehiveformat.
		//eventObject  *datamodels.EventForEsCaseObject = datamodels.NewEventForEsCaseObject()
		eventDetails *datamodels.EventCaseDetails     = datamodels.NewEventCaseDetails()

		eventObjectCustomFields  datamodels.CustomFields = datamodels.CustomFields{}
		eventDetailsCustomFields datamodels.CustomFields = datamodels.CustomFields{}
	)

	return vcase
}
