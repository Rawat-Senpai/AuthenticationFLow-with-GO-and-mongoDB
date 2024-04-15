package response

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// New Success Response creats a new success response
func SuccessResponse(data interface{}) Response {
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
