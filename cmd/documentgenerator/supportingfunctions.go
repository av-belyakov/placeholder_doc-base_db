package documentgenerator

import (
	"errors"
	"regexp"
	"slices"
)

// searchEventSource выполняет поиск источника события
func searchEventSource(fieldBranch string, a any) (string, bool) {
	if fieldBranch != "source" {
		return "", false
	}

	if v, ok := a.(string); ok {
		return v, true
	}

	return "", false
}

// getSensorIdFromDescription выполняет поиск идентификатора сенсора в поле description
func getSensorIdFromDescription(v string) (string, error) {
	rexSensorId := regexp.MustCompile(`СОА:\s-\s\*\*\x60(\d+)\x60\*\*`)
	tmp := rexSensorId.FindStringSubmatch(v)

	if len(tmp) <= 1 {
		return "", errors.New("there is no sensor ID in the accepted line")
	}

	return tmp[1], nil
}

// GetListIPAddr список ip адресов из элементов объекта
func GetListIPAddr(objects []IpAddressInformation) []string {
	newList := make([]string, 0, len(objects))

	for _, v := range objects {
		if slices.ContainsFunc(newList, func(elem string) bool {
			return elem == v.GetIpAddrString()
		}) {
			continue
		}

		newList = append(newList, v.GetIpAddrString())
	}

	return newList
}

// GetListSensorId список идентификаторов сенсоров
func GetListSensorId(objects []SensorInformation) []string {
	newList := make([]string, 0, len(objects))

	for _, v := range objects {
		if slices.ContainsFunc(newList, func(elem string) bool {
			return elem == v.GetSensorId()
		}) {
			continue
		}

		newList = append(newList, v.GetSensorId())
	}

	return newList
}
