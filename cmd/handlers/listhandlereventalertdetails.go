package handlers

import (
	"strings"

	alertdetailshf "github.com/av-belyakov/objectsthehiveformat/alertdetails"
	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

// NewListHandlerEventAlertDetails обработчик событий 'event.details.*' типа для объекта 'alert'
func NewListHandlerEventAlertDetails(details *alertdetailshf.EventAlertDetails) map[string][]func(any) {
	return map[string][]func(interface{}){
		"event.details.sourceRef":   {details.SetAnySourceRef},
		"event.details.title":       {details.SetAnyTitle},
		"event.details.description": {details.SetAnyDescription},
		"event.details.tags": {
			func(i any) {
				key, value := supportingfunctions.HandlerTag(i)
				if value == "" {
					return
				}

				value = strings.TrimSpace(value)
				value = strings.Trim(value, "\"")
				details.SetAnyTags(key, value)
			},
			details.SetAnyTagsAll,
		},
	}
}
