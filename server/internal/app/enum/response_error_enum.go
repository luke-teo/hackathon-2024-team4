package enum

type ErrorEnum struct {
	Code    string
	Message string
}

func ValidationFailedErrorEnum() ErrorEnum {
	return ErrorEnum{
		Code:    "general_68b329da",
		Message: "Validation failed",
	}
}

func InternalRequestHandlerErrorEnum() ErrorEnum {
	return ErrorEnum{
		Code:    "general_0YzHHt",
		Message: "",
	}
}

func InternalResponseHandlerErrorEnum() ErrorEnum {
	return ErrorEnum{
		Code:    "general_WrTRo9",
		Message: "",
	}
}
