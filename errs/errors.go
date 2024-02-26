package errs

type AppError struct {
	HTTPStatus int      `json:"httpStatus"`
	HTTPCode   string   `json:"httpCode"`
	RequestID  string   `json:"requestId,omitempty"`
	Errors     []Errors `json:"errors"`
}

type Errors struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func GenerateErrorResponse(httpStatus int, httpCode string, errors []Errors) *AppError {
	return &AppError{
		HTTPStatus: httpStatus,
		HTTPCode:   httpCode,
		Errors:     errors,
	}
}
