package service

import (
	"fmt"
	"live/utils"
	"live/module/model"
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

func CheckAuthValid(sessionId, accessToken string, vResult ApiResult) (userId int, valid bool) {
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

	gUserId, gErr := model.GetSession(sessionId)
	if gErr != nil {
		vResult.SetCode(API_SESSION_EXPIRED_ERROR)
		return
	}

	userId = gUserId
	valid = true
	return
}
