package model

import (
	"github.com/astaxie/beego/orm"
	"live/utils"
)

type User struct {
	Id     int
	Mobile string `orm:"column(mobile);size(100);unique"`
	Name   string `orm:"column(name);size(100);unique"`
	Pwd    string `orm:"column(pwd);size(32)"`
	Email  string `orm:"column(email);size(200);unique"`
}

//check mobile exists
func IsMobileExists(mobile string) (exists bool, err error) {
	qs := orm.NewOrm().QueryTable("User").Filter("Mobile", mobile).Limit(1)
	cnt, qErr := qs.Count()
	if qErr != nil {
		if qErr == orm.ErrNoRows {
			exists = false
		} else {
			err = qErr
		}
		return
	}
	if cnt == 1 {
		exists = true
	}
	return
}

func IsNameExists(name string) (exists bool, err error) {
	qs := orm.NewOrm().QueryTable("User").Filter("Name", name).Limit(1)
	cnt, qErr := qs.Count()
	if qErr != nil {
		if qErr == orm.ErrNoRows {
			exists = false
		} else {
			err = qErr
		}
		return
	}
	if cnt == 1 {
		exists = true
	}
	return
}

func IsEmailExists(email string) (exists bool, err error) {
	qs := orm.NewOrm().QueryTable("User").Filter("Email", email).Limit(1)
	cnt, qErr := qs.Count()
	if qErr != nil {
		if qErr == orm.ErrNoRows {
			exists = false
		} else {
			err = qErr
		}
		return
	}
	if cnt == 1 {
		exists = true
	}
	return
}

//get user
func GetUserByMobile(mobile string) (user *User, err error) {
	user = &User{}
	qs := orm.NewOrm().QueryTable("User").Filter("Mobile", mobile)
	err = qs.Limit(1).One(user)
	return
}

//create new user, if mobile exists, err it
func CreateNewUser(mobile, pwd, name, email string) (err error) {
	user := &User{
		Mobile: mobile,
		Name:   name,
		Pwd:    utils.Md5Hash(pwd),
		Email:  email,
	}
	_, err = orm.NewOrm().Insert(user)
	return
}
