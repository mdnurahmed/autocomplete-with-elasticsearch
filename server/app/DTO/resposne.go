package DTO

//ErrorResponse is response object when error occurs
type ErrorResponse struct {
	Name    string `json:"name" `
	Message string `json:"message"`
}

//NewErrorResponse returns a new ErrorResponse object
func NewErrorResponse(name string, msg string) *ErrorResponse {
	return &ErrorResponse{
		Name:    name,
		Message: msg,
	}
}
