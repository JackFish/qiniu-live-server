package model

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Session struct {
	Id        int
	UserId    int    `orm:"column(user_id)"`
	SessionId string `orm:"column(session_id)"`
}

func GetSession(sessionId string) (userId int, err error) {
	var session Session
	qs := orm.NewOrm().QueryTable("Session").Filter("SessionId", sessionId)
	qErr := qs.Limit(1).One(&session)
	if qErr != nil {
		err = errors.New(fmt.Sprintf("get session error, %s", qErr.Error()))
		return
	}
	userId = session.UserId

	return
}

func SetSession(userId int, sessionId string) (err error) {
	var session Session
	qs := orm.NewOrm().QueryTable("Session").Filter("UserId", userId)
	qErr := qs.Limit(1).One(&session)
	if qErr != nil {
		if qErr != orm.ErrNoRows {
			err = errors.New(fmt.Sprintf("set session error, %s", qErr.Error()))
			return
		}
	}
	if session.Id != 0 {
		session.SessionId = sessionId
		_, uErr := orm.NewOrm().Update(&session, "SessionId")
		if uErr != nil {
			err = errors.New(fmt.Sprintf("update session error, %s", uErr.Error()))
			return
		}
	} else {
		session.UserId = userId
		session.SessionId = sessionId
		_, cErr := orm.NewOrm().Insert(&session)
		if cErr != nil {
			err = errors.New(fmt.Sprintf("create session error, %s", cErr.Error()))
			return
		}
	}
	return
}
