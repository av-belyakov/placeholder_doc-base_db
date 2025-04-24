package handlers

import eventalerthf "github.com/av-belyakov/objectsthehiveformat/eventalert"

// NewListHandlerEventAlert обработчик событий типа 'event.*' для объекта 'alert'
func NewListHandlerEventAlert(event *eventalerthf.TypeEventForAlert) map[string][]func(any) {
	return map[string][]func(any){
		"event.base":           {event.SetAnyBase},
		"event.startDate":      {event.SetAnyStartDate},
		"event.rootId":         {event.SetAnyRootId},
		"event.objectId":       {event.SetAnyObjectId},
		"event.objectType":     {event.SetAnyObjectType},
		"event.organisation":   {event.SetAnyOrganisation},
		"event.organisationId": {event.SetAnyOrganisationId},
		"event.operation":      {event.SetAnyOperation},
		"event.requestId":      {event.SetAnyRequestId},
	}
}
