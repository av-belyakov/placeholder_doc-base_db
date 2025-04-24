package handlers

import (
	commonhf "github.com/av-belyakov/objectsthehiveformat/common"
)

// NewListHandlerAlertCustomFields обработчик событий 'alert.customFields.*' типа для объекта 'alert'
func NewListHandlerAlertCustomFields(customFields commonhf.CustomFields) map[string][]func(any) {
	return map[string][]func(any){
		//--------------- first-time ------------------
		"alert.customFields.first-time.order": {func(i any) {
			NewCustomFieldsElement("first-time", "date", &customFields)
			_, _, _, str := customFields["first-time"].Get()
			customFields["first-time"].Set(i, str)
		}},
		"alert.customFields.first-time.date": {func(i any) {
			NewCustomFieldsElement("first-time", "date", &customFields)
			_, order, _, _ := customFields["first-time"].Get()
			customFields["first-time"].Set(order, i)
		}},
		//--------------- last-time ------------------
		"alert.customFields.last-time.order": {func(i any) {
			NewCustomFieldsElement("last-time", "date", &customFields)
			_, _, _, str := customFields["last-time"].Get()
			customFields["last-time"].Set(i, str)
		}},
		"alert.customFields.last-time.date": {func(i any) {
			NewCustomFieldsElement("last-time", "date", &customFields)
			_, order, _, _ := customFields["last-time"].Get()
			customFields["last-time"].Set(order, i)
		}},
	}
}
