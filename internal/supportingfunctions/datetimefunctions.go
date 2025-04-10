package supportingfunctions

import "time"

// GetDateTimeFormatRFC3339 конвертирует числовое значение времени Unixtime
// в строку времени в формате RFC3339. Для корректной работы нужна дата в
// формате UnixMilli-секунд (13 символов)
func GetDateTimeFormatRFC3339(dt int64) string {
	return time.UnixMilli(dt).Format(time.RFC3339)
}
