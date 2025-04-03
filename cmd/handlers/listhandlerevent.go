package handlers

import eventcasehf "github.com/av-belyakov/objectsthehiveformat/eventcase"

// NewListHandlerEventCase обработчик событий типа 'case' для 'event'
func NewListHandlerEventCase(event *eventcasehf.TypeEventForCase) map[string][]func(interface{}) {
	return map[string][]func(any){
		"event.rootId":         {event.SetAnyRootId},
		"event.objectId":       {event.SetAnyObjectId},
		"event.objectType":     {event.SetAnyObjectType},
		"event.base":           {event.SetAnyBase},
		"event.startDate":      {event.SetAnyStartDate},
		"event.requestId":      {event.SetAnyRequestId},
		"event.organisation":   {event.SetAnyOrganisation},
		"event.organisationId": {event.SetAnyOrganisationId},
		"event.operation":      {event.SetAnyOperation},
	}
}
