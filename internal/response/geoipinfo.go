package response

// ******* для ответов на запрос о географическом расположении ip адресов *******
// ResponseGeoIpInformation ответ от внешнего сервиса по географическому
// местоположению ip адресов
type ResponseGeoIpInformation struct {
	Informations []GeoIpInformation `json:"found_information"`
	Source       string             `json:"source"`
	TaskId       string             `json:"task_id"`
	Error        string             `json:"error"`
}

// GeoIpInformation плодробная информация по географическому местоположению ip адресов
type GeoIpInformation struct {
	IpAddr      string         `json:"ip_address"`
	Code        string         `json:"code"`
	Country     string         `json:"country"`
	City        string         `json:"city"`
	Subnet      string         `json:"subnet"`
	UpdatedAt   string         `json:"updated_at"`
	Error       string         `json:"error"`
	RangeIpAddr RangeIpAddress `json:"ip_range"`
}

type RangeIpAddress struct {
	IpFirst string `json:"ip_first"`
	IpLast  string `json:"ip_last"`
}
