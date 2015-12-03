package model

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"live/utils"
	"time"
)

type LiveVideo struct {
	Id          int
	User        *User     `orm:"column(user_id);rel(fk)"`
	PublishId   string    `orm:"column(publish_id);size(32)"`
	Title       string    `orm:"column(title);size(100)"`
	StreamId    string    `orm:"column(stream_id);size(32)"`
	StartTime   int64     `orm:"column(start_time)"`
	EndTime     int64     `orm:"column(end_time)"`
	Quality     int       `orm:"column(quality)"`
	Orientation int       `orm:"column(orientation)"`
	CreateTime  time.Time `orm:"column(create_time);auto_now_add"`
}

//@param streamId
func IsStreamPublishing(streamId string) (publishing bool, err error) {
	qs := orm.NewOrm().QueryTable("LiveVideo").Filter("StreamId", streamId).Filter("EndTime", 0)
	cnt, qErr := qs.Count()
	if qErr != nil {
		err = qErr
		return
	}
	publishing = (cnt > 0)
	return
}

//@param streamId
//@param streamTitle
//@return publishId if err==nil and pubilshId==nil, not allow taken
//streamId can only be taken once a time
func CreateNewPublish(userId int, streamId, streamTitle string, streamQuality, streamOrientation int) (publishId string, err error) {
	timestamp := time.Now().Unix()
	orm.NewOrm().QueryTable("LiveVideo").Filter("StreamId", streamId).Filter("EndTime", 0).Update(orm.Params{
		"EndTime": timestamp,
	})

	publishId = utils.Md5Hash(fmt.Sprintf("%s:%s:%d", streamId, streamTitle, timestamp))
	user := User{Id: userId}
	newPublish := LiveVideo{
		User:        &user,
		PublishId:   publishId,
		Title:       streamTitle,
		StreamId:    streamId,
		Quality:     streamQuality,
		Orientation: streamOrientation,
		StartTime:   timestamp,
	}

	_, cErr := orm.NewOrm().Insert(&newPublish)
	if cErr != nil {
		err = cErr
		return
	}

	return
}

//@param publishId
func CompletePublish(publishId string) (err error) {
	_, uErr := orm.NewOrm().QueryTable("LiveVideo").Filter("PublishId", publishId).Update(orm.Params{
		"EndTime": time.Now().Unix(),
	})
	if uErr != nil {
		err = uErr
		return
	}
	return
}

//@return videoList
func GetLiveVideoList(videoList *[]LiveVideo) (err error) {
	qs := orm.NewOrm().QueryTable("LiveVideo").Exclude("EndTime", 0).RelatedSel()
	_, err = qs.Limit(100).OrderBy("-CreateTime").All(videoList)
	return
}

//@return videoList
func GetLiveStreamList(videoList *[]LiveVideo) (err error) {
	qs := orm.NewOrm().QueryTable("LiveVideo").Filter("EndTime", 0).RelatedSel()
	_, err = qs.Limit(100).OrderBy("-CreateTime").All(videoList)
	return
}

//@return videoList
func GetMyLiveVideoList(userId int, videoList *[]LiveVideo) (err error) {
	qs := orm.NewOrm().QueryTable("LiveVideo").Filter("User__id", userId).Exclude("EndTime", 0).RelatedSel()
	_, err = qs.Limit(100).OrderBy("-CreateTime").All(videoList)
	return
}

func GetLiveVideoByPublishId(publishId string, liveVideo *LiveVideo) (err error) {
	qs := orm.NewOrm().QueryTable("LiveVideo").Filter("PublishId", publishId).Limit(1)
	qErr := qs.One(liveVideo)
	if qErr != nil {
		if qErr != orm.ErrNoRows {
			err = qErr
		}
	}

	return
}
