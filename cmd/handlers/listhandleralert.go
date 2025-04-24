package handlers

import (
	alerttemplate "github.com/av-belyakov/objectsthehiveformat/alert"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// NewListHandlerAlert обработчик событий 'alert.*' типа для объекта 'alert'
func NewListHandlerAlert(alert *alerttemplate.TypeAlert) map[string][]func(any) {
	return map[string][]func(any){
		"alert.follow":       {alert.SetAnyFollow},
		"alert.severity":     {alert.SetAnySeverity},
		"alert.tlp":          {alert.SetAnyTlp},
		"alert.date":         {alert.SetAnyDate},
		"alert.createdAt":    {alert.SetAnyCreatedAt},
		"alert.updatedAt":    {alert.SetAnyUpdatedAt},
		"alert.createdBy":    {alert.SetAnyCreatedBy},
		"alert.updatedBy":    {alert.SetAnyUpdatedBy},
		"alert.source":       {alert.SetAnySource},
		"alert.title":        {alert.SetAnyTitle},
		"alert.id":           {alert.SetAnyId},
		"alert._id":          {alert.SetAnyUnderliningId},
		"alert.status":       {alert.SetAnyStatus},
		"alert.type":         {alert.SetAnyType},
		"alert._type":        {alert.SetAnyUnderliningType},
		"alert.description":  {alert.SetAnyDescription},
		"alert.caseTemplate": {alert.SetAnyCaseTemplate},
		"alert.sourceRef":    {alert.SetAnySourceRef},
		"alert.tags": {
			func(i any) {
				key, value := supportingfunctions.HandlerTag(i)
				if value == "" {
					return
				}

				alert.SetAnyTags(key, value)
			},
			alert.SetAnyTagsAll},
	}
}
