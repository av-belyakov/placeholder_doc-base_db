package documentgenerator

import (
	"errors"
	"regexp"
)

// searchEventSource выполняет поиск источника события
func searchEventSource(fieldBranch string, value any) (string, bool) {
	if fieldBranch != "source" {
		return "", false
	}

	if v, ok := value.(string); ok {
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
