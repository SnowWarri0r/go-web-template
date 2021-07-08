package util

type ResultStatus struct {
	status  int
	message string
}

var (
	SUCCESS = ResultStatus{status: 0, message: "success"}
	FAILURE = ResultStatus{status: 1, message: "failure"}
	WRONG_PARAMETERS = ResultStatus{status: 2, message: "wrong param"}
	USELESS_REQUEST = ResultStatus{status: 3, message: "useless request"}
	INTERNAL_SERVER_ERROR = ResultStatus{status: 500, message: "Internal Server Error"}
)

func (r ResultStatus) SetTheMessage(message string) ResultStatus {
	r.message =message
	return r
}