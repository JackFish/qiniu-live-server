package server

import (
	"live/module/service"
	"net/http"
	"strings"
)

//login
//@param mobile
//@param pwd
func (this *LiveServer) serveLogin(w http.ResponseWriter, req *http.Request) {
	loginResult := service.LoginResult{}

	req.ParseForm()

	mobile := strings.TrimSpace(req.FormValue("mobile"))
	pwd := strings.TrimSpace(req.FormValue("pwd"))

	service.UserLogin(mobile, pwd, &loginResult)
	this.serveResultJson(w, req, http.StatusOK, &loginResult)
}

//signup
//@param mobile
//@param pwd
func (this *LiveServer) serveSignup(w http.ResponseWriter, req *http.Request) {
	signupResult := service.SignupResult{}
	req.ParseForm()

	mobile := strings.TrimSpace(req.FormValue("mobile"))
	pwd := strings.TrimSpace(req.FormValue("pwd"))
	name := strings.TrimSpace(req.FormValue("name"))
	email := strings.TrimSpace(req.FormValue("email"))

	service.UserSignup(mobile, pwd, name, email, &signupResult)
	this.serveResultJson(w, req, http.StatusOK, &signupResult)
}
