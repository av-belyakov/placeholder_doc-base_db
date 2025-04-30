package decoderjsondocuments

import (
	"fmt"
	"reflect"

	"github.com/av-belyakov/placeholder_doc-base_db/interfaces"
)

func processingReflectAnySimpleType(chInput chan<- interfaces.CustomJsonDecoder, name any, anyType any, fieldBranch string) any {
	var nameStr string
	r := reflect.TypeOf(anyType)

	if n, ok := name.(int); ok {
		nameStr = fmt.Sprint(n)
	} else if n, ok := name.(string); ok {
		nameStr = n
	}

	if r == nil {
		return anyType
	}

	switch r.Kind() {
	case reflect.String:
		result := reflect.ValueOf(anyType).String()
		chInput <- &ChanInputSettings{
			FieldName:   nameStr,
			ValueType:   "string",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		result := reflect.ValueOf(anyType).Int()
		chInput <- &ChanInputSettings{
			FieldName:   nameStr,
			ValueType:   "int",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result := reflect.ValueOf(anyType).Uint()
		chInput <- &ChanInputSettings{
			FieldName:   nameStr,
			ValueType:   "uint",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Float32, reflect.Float64:
		result := reflect.ValueOf(anyType).Float()
		chInput <- &ChanInputSettings{
			FieldName:   nameStr,
			ValueType:   "float",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Bool:
		result := reflect.ValueOf(anyType).Bool()
		chInput <- &ChanInputSettings{
			FieldName:   nameStr,
			ValueType:   "bool",
			Value:       result,
			FieldBranch: fieldBranch,
		}

		return result
	}

	return anyType
}

func processingReflectMap(chInput chan<- interfaces.CustomJsonDecoder, l map[string]any, fieldBranch string) map[string]any {
	var (
		newMap  map[string]any
		newList []any
	)
	nl := map[string]any{}

	for k, v := range l {
		var fbTmp string
		r := reflect.TypeOf(v)

		if r == nil {
			continue
		}

		fbTmp = fieldBranch
		if fbTmp == "" {
			fbTmp += k
		} else {
			fbTmp += "." + k
		}

		switch r.Kind() {
		case reflect.Map:
			if v, ok := v.(map[string]any); ok {
				newMap = processingReflectMap(chInput, v, fbTmp)
				nl[k] = newMap
			}

		case reflect.Slice:
			if v, ok := v.([]any); ok {
				newList = processingReflectSlice(chInput, v, fbTmp)
				nl[k] = newList
			}

		default:
			nl[k] = processingReflectAnySimpleType(chInput, k, v, fbTmp)
		}
	}

	return nl
}

func processingReflectSlice(chInput chan<- interfaces.CustomJsonDecoder, l []any, fieldBranch string) []any {
	var (
		newMap  map[string]any
		newList []any
	)
	nl := make([]any, 0, len(l))

	for k, v := range l {
		r := reflect.TypeOf(v)

		if r == nil {
			continue
		}

		switch r.Kind() {
		case reflect.Map:
			if v, ok := v.(map[string]any); ok {
				newMap = processingReflectMap(chInput, v, fieldBranch)

				nl = append(nl, newMap)
			}

		case reflect.Slice:
			if v, ok := v.([]any); ok {
				newList = processingReflectSlice(chInput, v, fieldBranch)

				nl = append(nl, newList...)
			}

		default:
			nl = append(nl, processingReflectAnySimpleType(chInput, k, v, fieldBranch))
		}
	}

	return nl
}
