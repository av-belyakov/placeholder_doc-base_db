package handlers

import (
	alertobjecthf "github.com/av-belyakov/objectsthehiveformat/alertobject"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// NewListHandlerEventAlertObject обработчик событий типа 'event.object.*' для объекта 'alert'
func NewListHandlerEventAlertObject(object *alertobjecthf.EventAlertObject) map[string][]func(any) {
	return map[string][]func(any){
		"event.object.tlp":          {object.SetAnyTlp},
		"event.object._id":          {object.SetAnyUnderliningId},
		"event.object.id":           {object.SetAnyId},
		"event.object.createdBy":    {object.SetAnyCreatedBy},
		"event.object.updatedBy":    {object.SetAnyUpdatedBy},
		"event.object.createdAt":    {object.SetAnyCreatedAt},
		"event.object.updatedAt":    {object.SetAnyUpdatedAt},
		"event.object._type":        {object.SetAnyUnderliningType},
		"event.object.title":        {object.SetAnyTitle},
		"event.object.description":  {object.SetAnyDescription},
		"event.object.status":       {object.SetAnyStatus},
		"event.object.date":         {object.SetAnyDate},
		"event.object.type":         {object.SetAnyType},
		"event.object.objectType":   {object.SetAnyObjectType},
		"event.object.source":       {object.SetAnySource},
		"event.object.sourceRef":    {object.SetAnySourceRef},
		"event.object.case":         {object.SetAnyCase},
		"event.object.caseTemplate": {object.SetAnyCaseTemplate},
		"event.object.tags": {
			func(i any) {
				key, value := supportingfunctions.HandlerTag(i)
				if value == "" {
					return
				}

				object.SetAnyTags(key, value)
			},
			object.SetAnyTagsAll},
	}
}
