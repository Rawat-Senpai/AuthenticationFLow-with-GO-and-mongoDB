package response

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//  NewErrorResponse creates a new error response

func NewErrorResponse(success bool, message string) Response {
	return Response{
		Success: success,
		Message: message,
	}
}

// New Success Response creats a new success response

func NewSuccessResponse(data interface{}) Response {
	return Response{
		Success: true,
		Message: "Success",
		Data:    data,
	}
}

// Certain Error
func ErrorResponse(message string) Response {
	return Response{
		Success: false,
		Message: message,
	}
}
