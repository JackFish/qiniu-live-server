package service

import (
	"github.com/astaxie/beego/orm"
	"live/module/model"
	"live/utils"
)

type LoginResult struct {
	ApiResult
	UserName  string `json:"userName,omitempty"`
	SessionId string `json:"sessionId,omitempty"`
}

func (this *LoginResult) SetOk() {
	this.Code = API_OK
	this.Desc = "login success"
}

type SignupResult struct {
	ApiResult
}

func (this *SignupResult) SetOk() {
	this.Code = API_OK
	this.Desc = "signup success"
}

//@param mobile
//@param pwd
//@output loginResult
//@return ok, not found, error
func UserLogin(mobile, pwd string, loginResult *LoginResult) {
	user, qErr := model.GetUserByMobile(mobile)
	if qErr != nil {
		if qErr == orm.ErrNoRows {
			loginResult.SetCode(API_USER_NOT_FOUND_ERROR)
		} else {
			loginResult.SetCode(API_SERVER_ERROR)
		}
		return
	}

	if user.Pwd != utils.Md5Hash(pwd) {
		loginResult.SetCode(API_USER_PWD_ERROR)
		return
	}

	sessionId := utils.CreateSessionId(user.Mobile, user.Pwd)
	cErr := model.SetSession(user.Id, sessionId)
	if cErr != nil {
		loginResult.SetCode(API_SERVER_ERROR)
		return
	}

	loginResult.UserName = user.Name
	loginResult.SessionId = sessionId
	loginResult.SetOk()

	return
}

//@param mobile
//@param pwd
//@output signupResult
//@return ok, user exists, error
func UserSignup(mobile, pwd, name, email string, signupResult *SignupResult) {
	if mobile == "" {
		signupResult.SetFormatCode(API_PARAM_ERROR, "mobile is empty")
		return
	}

	if pwd == "" {
		signupResult.SetFormatCode(API_PARAM_ERROR, "pwd is empty")
		return
	}

	if name == "" {
		signupResult.SetFormatCode(API_PARAM_ERROR, "name is empty")
		return
	}

	if email == "" {
		signupResult.SetFormatCode(API_PARAM_ERROR, "email is empty")
		return
	}

	cErr := model.CreateNewUser(mobile, pwd, name, email)
	if cErr != nil {
		if exists, qErr := model.IsMobileExists(mobile); qErr == nil {
			if exists {
				signupResult.SetCode(API_PHONE_EXISTS_ERROR)
			} else {
				if exists, qErr = model.IsNameExists(name); qErr == nil {
					if exists {
						signupResult.SetCode(API_NAME_EXISTS_ERROR)
					} else {
						if exists, qErr = model.IsEmailExists(email); qErr == nil {
							signupResult.SetCode(API_EMAIL_EXISTS_ERROR)
						} else {
							signupResult.SetCode(API_SERVER_ERROR)
						}
					}
				} else {
					signupResult.SetCode(API_SERVER_ERROR)
				}
			}
		} else {
			signupResult.SetCode(API_SERVER_ERROR)
		}
		return
	}

	signupResult.SetOk()
	return

}
