package response

// Standart Response Struct
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

// Func to Build a Successfull Response
func BuildSuccessResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Errors:  nil,
		Data:    data,
	}
}

// Func to Build a Failed Response
func BuildFailedResponse(message string, errors interface{}) Response {
	return Response{
		Success: false,
		Message: message,
		Errors:  errors,
		Data:    nil,
	}
}
