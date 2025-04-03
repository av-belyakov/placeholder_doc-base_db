package handlers

import (
	"fmt"
	"strings"

	commonhf "github.com/av-belyakov/objectsthehiveformat/common"
)

// NewCustomFieldsElement данный метод, на основании значения objType
// определяет ссылку на какой пользовательский тип datamodels.CustomFieldStringType,
// datamodels.CustomFieldDateType или datamodels.CustomFieldIntegerType добавить
// добавить к map[string]datamodels.CustomerFields. Это сделанно для того что бы
// была возможность подобрать польховательский тип на основе типов свойств.
// Например, нужен тип с полями order int и string типа string, а в другой раз
// тип с полями date string и order int.
// func newCustomFieldsElement(elem, objType string, customFields *map[string]datamodels.CustomerFields) {
func NewCustomFieldsElement(elem, objType string, customFields *commonhf.CustomFields) {
	if _, ok := (*customFields)[elem]; !ok {
		switch objType {
		case "string":
			(*customFields)[elem] = &commonhf.CustomFieldStringType{}
		case "date":
			(*customFields)[elem] = &commonhf.CustomFieldDateType{
				Date: "1970-01-01T03:00:00+03:00",
			}
		case "integer":
			(*customFields)[elem] = &commonhf.CustomFieldIntegerType{}
		case "boolen":
			(*customFields)[elem] = &commonhf.CustomFieldBoolenType{}
		default:
			(*customFields)[elem] = &commonhf.CustomFieldStringType{}
		}
	}
}

// HandlerTag выполняет обработку тегов, разделяя тег на его тип и значение
func HandlerTag(i any) (key, value string) {
	isExistValidTag := func(item string) bool {
		validListTags := []string{
			"geo",
			"geoip",
			"reason",
			"sensor",
			"misp",
			"ioc",
		}

		for _, v := range validListTags {
			if strings.Contains(item, v) {
				return true
			}
		}

		return false
	}

	tag := strings.ToLower(fmt.Sprint(i))

	if isExistValidTag(tag) && strings.Contains(tag, "=") {
		elements := strings.Split(tag, "=")
		if len(elements) > 1 {
			if strings.Contains(elements[0], "geo") {
				return elements[0], strings.ToUpper(elements[1])
			}

			return elements[0], elements[1]
		}
	}

	return tag, ""
}
