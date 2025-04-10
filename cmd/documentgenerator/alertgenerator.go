package documentgenerator

import (
	alertdetails "github.com/av-belyakov/objectsthehiveformat/alertdetails"
	alertobjects "github.com/av-belyakov/objectsthehiveformat/alertobject"
	eventalert "github.com/av-belyakov/objectsthehiveformat/eventalert"

	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
)

func AlertGeneratorchInput(<-chan interfaces.CustomJsonDecoder) (string, *VerifiedAlert, map[string]string) {
	var (
		rootId string
		// список не обработанных полей
		listRawFields map[string]string = make(map[string]string)

		//Финальный объект
		verifiedAlert *VerifiedAlert = NewVerifiedAlert()

		event        *eventalert.TypeEventForAlert   = eventalert.NewTypeEventForAlert()
		eventObjects *alertobjects.EventAlertObject  = alertobjects.NewEventAlertObject()
		eventDetails *alertdetails.EventAlertDetails = alertdetails.NewEventAlertDetails()
	)

	return rootId, verifiedAlert, listRawFields
}
