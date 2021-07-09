package util

type Result struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func New(resultStatus ResultStatus, data interface{}) Result {
	return Result{
		Status:  resultStatus.status,
		Message: resultStatus.message,
		Data:    data,
	}
}

func Success() Result {
	return New(SUCCESS, nil)
}

func SuccessInfo(message string) Result {
	return New(SUCCESS.SetTheMessage(message), nil)
}

func SuccessInfoAndData(message string, data interface{}) Result {
	return New(SUCCESS.SetTheMessage(message), data)
}

func SuccessData(data interface{}) Result {
	return New(SUCCESS, data)
}

func Failure() Result {
	return New(FAILURE, nil)
}

func FailureInfo(message string) Result {
	return New(FAILURE.SetTheMessage(message), nil)
}

func FailureInfoAndData(message string, data interface{}) Result {
	return New(FAILURE.SetTheMessage(message), data)
}

func FailureData(data interface{}) Result {
	return New(FAILURE, data)
}

func WrongParam() Result {
	return New(WrongParameters, nil)
}

func WrongParamInfo(message string) Result {
	return New(WrongParameters.SetTheMessage(message), nil)
}

func WrongParamInfoAndData(message string, data interface{}) Result {
	return New(WrongParameters.SetTheMessage(message), data)
}

func WrongParamData(data interface{}) Result {
	return New(WrongParameters, data)
}

func UselessReq() Result {
	return New(UselessRequest, nil)
}

func UselessReqInfo(message string) Result {
	return New(UselessRequest.SetTheMessage(message), nil)
}

func UselessReqInfoAndData(message string, data interface{}) Result {
	return New(UselessRequest.SetTheMessage(message), data)
}

func UselessReqData(data interface{}) Result {
	return New(UselessRequest, data)
}

func Error() Result {
	return New(InternalServerError, nil)
}
