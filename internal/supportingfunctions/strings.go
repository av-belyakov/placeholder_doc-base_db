package supportingfunctions

import (
	"fmt"
	"strings"
)

// GetWhitespace возвращает необходимое количество пробелов
func GetWhitespace(num int) string {
	var str string

	if num == 0 {
		return str
	}

	for i := 0; i < num; i++ {
		str += "  "
	}

	return str
}

// JoinRawFieldsToString объединяет список необработанных полей в строку
func JoinRawFieldsToString(list map[string]string, tag, id string) string {
	var str strings.Builder = strings.Builder{}

	for k, v := range list {
		str.WriteString(fmt.Sprintf("\n\t%s %s field: '%s', value: '%s'", tag, id, k, v))
	}

	return str.String()
}
