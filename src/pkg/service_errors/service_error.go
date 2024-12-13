package service_errors

type ServiceError struct {
	EndUserMessage   string `json:"endUserMessage"`
	TechnicalMessage string `json:"technicalMessage"`
	Err              error
}

func (service *ServiceError) Error() string {
	return service.EndUserMessage
}
