package model

import (
	"errors"
	"github.com/astaxie/beego/orm"
)

type LiveStream struct {
	Id       int
	UserId   int    `orm:"column(user_id)"`
	StreamId string `orm:"column(stream_id);size(32)"`
}

//@param user id
//@return err if exec error
//@return streamId not empty or empty(stands for none)
func GetStreamIdOfUser(userId int) (streamId string, err error) {
	var liveStream LiveStream
	qs := orm.NewOrm().QueryTable("LiveStream").Filter("UserId", userId)
	qErr := qs.Limit(1).One(&liveStream)
	if qErr != nil {
		if qErr != orm.ErrNoRows {
			err = errors.New("get stream of user error")
		}
		return
	}

	streamId = liveStream.StreamId

	return
}

//@param user id
//@param stream id
//@return exec error
func SetStreamIdOfUser(userId int, streamId string) (err error) {
	liveStream := LiveStream{
		UserId:   userId,
		StreamId: streamId,
	}

	_, cErr := orm.NewOrm().Insert(&liveStream)
	if cErr != nil {
		err = errors.New("set stream of user error")
		return
	}
	return
}
