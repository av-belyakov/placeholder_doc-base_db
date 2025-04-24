package supportingfunctions

import (
	"fmt"
	"strings"
)

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
