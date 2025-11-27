// internal/models/legacy/service_get.go
package legacy

type ServiceGetResponse struct {
	Data ServiceGetData `json:"data"`
}

type ServiceGetData struct {
	Items []ServiceGet `json:"items"`
	Total int          `json:"total"`
}

type ServiceGet struct {
	ServiceVPS
	// В будущем можно добавить другие embedded структуры для разных типов услуг
	// ServiceDedicated для выделенных серверов
}
