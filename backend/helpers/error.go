package helpers

func ErrorResponse(errorMessage string) map[string]string {
	message := make(map[string]string)
	message["message"] = errorMessage
	return message
}

func RandomErrorResponse() map[string]string {
	return ErrorResponse("something went wrong")
}


func InvaildBodyErrorResponse()map[string]string{
	return ErrorResponse("invaild body")
}