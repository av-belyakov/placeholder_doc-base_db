package supportingfunctions

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"reflect"
	"sync"
)

// ElementsFromJSON элементы полученные при обработки JSON объекта
type ElementsFromJSON struct {
	mutex sync.RWMutex
	Data  map[string]Element
}

// CommonValues общие, для типов, значения
type CommonValues struct {
	Value     any    //любые передаваемые данные
	FieldName string //наименование поля
	ValueType string //тип передаваемого значения (string, int и т.д.)
}

// Element описание объекта
type Element struct {
	CommonValues
}

type chResult struct {
	CommonValues
	FieldBranch string //'путь' до значения в как в JSON формате, например 'event.details.customFields.class'
}

// GetElementsFromJSON читает JSON объект методом рефлексии, возвращает результат,
// содержащий перечень путей до элемента и значение элемента с указанием его имени
// и типа
func GetElementsFromJSON(ctx context.Context, data []byte) (map[string]Element, error) {
	ctxCancel, cancel := context.WithCancel(ctx)

	chRes := make(chan chResult)
	defer close(chRes)

	result := ElementsFromJSON{Data: map[string]Element{}}
	copyResult := map[string]Element{}

	//обработчик входящих данных
	go func(ctx context.Context, rst *ElementsFromJSON, ch <-chan chResult) {
		rst.mutex.Lock()
		defer rst.mutex.Unlock()

		for {
			select {
			case <-ctx.Done():
				return

			case data, ok := <-ch:
				if !ok {
					return
				}

				rst.Data[data.FieldBranch] = Element{
					CommonValues{
						FieldName: data.FieldName,
						ValueType: data.ValueType,
						Value:     data.Value,
					},
				}
			}
		}
	}(ctxCancel, &result, chRes)

	listMap := map[string]any{}
	if err := json.Unmarshal(data, &listMap); err == nil {
		//для карт
		if len(listMap) == 0 {
			cancel()

			return copyResult, errors.New("error decoding the json file, it may be empty")
		}

		_ = processingReflectMap(chRes, listMap, "")

	} else {
		// для срезов
		listSlice := []any{}
		if err = json.Unmarshal(data, &listSlice); err != nil {
			cancel()

			return copyResult, err
		}

		if len(listSlice) == 0 {
			cancel()

			return copyResult, errors.New("error decoding the json message, it may be empty")
		}

		_ = processingReflectSlice(chRes, listSlice, "")
	}

	cancel()

	result.mutex.RLock()
	maps.Copy(copyResult, result.Data)
	result.mutex.Unlock()

	return copyResult, nil
}

func processingReflectMap(ch chan<- chResult, list map[string]any, fieldBranch string) map[string]any {
	var (
		newMap  map[string]any
		newList []any

		nl map[string]any = map[string]any{}
	)

	for k, v := range list {
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
				newMap = processingReflectMap(ch, v, fbTmp)
				nl[k] = newMap
			}

		case reflect.Slice:
			if v, ok := v.([]any); ok {
				newList = processingReflectSlice(ch, v, fbTmp)
				nl[k] = newList
			}

		default:
			nl[k] = processingReflectAnySimpleType(ch, k, v, fbTmp)
		}
	}

	return nl
}

func processingReflectSlice(ch chan<- chResult, list []any, fieldBranch string) []any {
	var (
		newMap  map[string]any
		newList []any

		nl []any = make([]any, 0, len(list))
	)

	for k, v := range list {
		r := reflect.TypeOf(v)

		if r == nil {
			continue
		}

		switch r.Kind() {
		case reflect.Map:
			if v, ok := v.(map[string]any); ok {
				newMap = processingReflectMap(ch, v, fieldBranch)

				nl = append(nl, newMap)
			}

		case reflect.Slice:
			if v, ok := v.([]any); ok {
				newList = processingReflectSlice(ch, v, fieldBranch)

				nl = append(nl, newList...)
			}

		default:
			nl = append(nl, processingReflectAnySimpleType(ch, k, v, fieldBranch))
		}
	}

	return nl
}

func processingReflectAnySimpleType(ch chan<- chResult, name any, anyType any, fieldBranch string) any {
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
		ch <- chResult{
			CommonValues: CommonValues{
				FieldName: nameStr,
				ValueType: "string",
				Value:     result,
			},
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
		result := reflect.ValueOf(anyType).Int()
		ch <- chResult{
			CommonValues: CommonValues{
				FieldName: nameStr,
				ValueType: "int",
				Value:     result,
			},
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		result := reflect.ValueOf(anyType).Uint()
		ch <- chResult{
			CommonValues: CommonValues{
				FieldName: nameStr,
				ValueType: "uint",
				Value:     result,
			},
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Float32, reflect.Float64:
		result := reflect.ValueOf(anyType).Float()
		ch <- chResult{
			CommonValues: CommonValues{
				FieldName: nameStr,
				ValueType: "float",
				Value:     result,
			},
			FieldBranch: fieldBranch,
		}

		return result

	case reflect.Bool:
		result := reflect.ValueOf(anyType).Bool()
		ch <- chResult{
			CommonValues: CommonValues{
				FieldName: nameStr,
				ValueType: "bool",
				Value:     result,
			},
			FieldBranch: fieldBranch,
		}

		return result
	}

	return anyType
}
