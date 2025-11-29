package legacy

// ServiceProlongRequest - запрос на продление услуги через Legacy API
type ServiceProlongRequest struct {
	Method string `json:"method"`
	Term   string `json:"term"`
	Count  int64  `json:"count"`
}

// ServiceProlongResponse - ответ на запрос продления услуги
type ServiceProlongResponse struct {
	Data struct {
		Transaction Transaction `json:"transaction"`
	} `json:"data"`
}
