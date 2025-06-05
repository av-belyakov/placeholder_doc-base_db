package request

// RequestGeoIP запрос к модулю enricher_geoip для получения информации по ip
type RequestGeoIP struct {
	Source          string   `json:"source"`
	TaskId          string   `json:"task_id"`
	ListIpAddresses []string `json:"list_ip_addresses"`
}

// RequestSensorInformation запрос к модулю enricher_sensor_information для
// получения информации по сенсорам
type RequestSensorInformation struct {
	Source      string   `json:"source"`
	TaskId      string   `json:"task_id"`
	ListSensors []string `json:"list_sensors"`
}
