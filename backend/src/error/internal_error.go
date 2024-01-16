package error

type InternalError struct {
	EndUserMessage   string `json:"endUserMessage"`
	TechnicalMessage string `json:"technicalMessage"`
	Err              error
}

func (error *InternalError) Error() string {
	return error.EndUserMessage
}
