package http_error

import "github.com/uccu/go-stringify"

type ResData map[string]interface{}

type HttpError struct {
	ErrorCode int    `json:"code" example:"400"`
	ErrorMsg  string `json:"msg" example:"status bad request"`
}

func HttpErrorEx(e HttpError, i interface{}) HttpError {
	return HttpError{
		ErrorCode: e.ErrorCode,
		ErrorMsg:  e.ErrorMsg + ": " + stringify.ToString(i),
	}
}

var MissingParametersError = HttpError{
	ErrorCode: 501,
	ErrorMsg:  "缺少参数",
}

var WrongPage = HttpError{
	ErrorCode: 502,
	ErrorMsg:  "wrong page",
}

var NoXGitlabToken = HttpError{
	ErrorCode: 1001,
	ErrorMsg:  "1001",
}

var XGitlabTokenNotMatch = HttpError{
	ErrorCode: 1002,
	ErrorMsg:  "1002",
}
