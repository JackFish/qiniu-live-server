package service

import (
	"fmt"
	"live/utils"
)

type ApiResult struct {
	Code int    `json:"code"`
	Desc string `json:"desc"`
}

func (this *ApiResult) SetCode(code int) {
	this.Code = code
	this.Desc = ApiStatus[this.Code]
}

func (this *ApiResult) SetFormatCode(code int, val string) {
	this.Code = code
	this.Desc = fmt.Sprintf(ApiStatus[this.Code], val)
}

func CheckAuthValid(sessionId, accessToken string, vResult ApiResult) (valid bool) {
	if sessionId == "" {
		vResult.SetFormatCode(API_PARAM_ERROR, "session id is empty")
		return
	}

	if accessToken == "" {
		vResult.SetFormatCode(API_PARAM_ERROR, "access token is empty")
		return
	}

	if !utils.IsAccessTokenValid(sessionId, accessToken) {
		vResult.SetCode(API_UNAUTHORIZED_ERROR)
		return
	}

	valid = true
	return
}
