package helper

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

type EmptyObject struct {
}

func BuildSuccessResponse(message string, data interface{}) Response {
	successResponse := Response{
		Success: true,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
	return successResponse
}

func BuildFailedResponse(message string, errors interface{}) Response {
	failedResponse := Response{
		Success: false,
		Message: message,
		Data:    nil,
		Errors:  errors,
	}
	return failedResponse
}
