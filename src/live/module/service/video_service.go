package service

import (
	"live/module/model"
	"live/module/service/pilis"
	"github.com/qiniu/log"
)

type PlaybackVideo struct {
	User       string `json:"user"`
	Title      string `json:"title"`
	PublishId  string `json:"publishId"`
	CreateTime int64  `json:"createTime"`
}

type PlaybackVideoListResult struct {
	ApiResult
	VideoList []PlaybackVideo `json:"videoList,omitempty"`
}

func (this *PlaybackVideoListResult) SetOk() {
	this.Code = API_OK
	this.Desc = "get video list success"
}

//////////////////

type PublishingVideo struct {
	User       string `json:"user"`
	Title      string `json:"title"`
	PublishId  string `json:"publishId"`
	CreateTime int64  `json:"createTime"`
}

type PublishingVideoListResult struct {
	ApiResult
	VideoList []PublishingVideo `json:"videoList"`
}

func (this *PublishingVideoListResult) SetOk() {
	this.Code = API_OK
	this.Desc = "get publishing list success"
}

//////////////////////

type VideoPlayResult struct {
	ApiResult
	Orientation int `json:"orientation"`
	PlayUrls map[string]string `json:"playUrls"`
}

func (this *VideoPlayResult) SetOk() {
	this.Code = API_OK
	this.Desc = "get video play url success"
}

//////////////////////

type StreamPlayResult struct {
	ApiResult
	Orientation int `json:"orientation"`
	PlayUrls map[string]string `json:"playUrls"`
}

func (this *StreamPlayResult) SetOk() {
	this.Code = API_OK
	this.Desc = "get stream play url success"
}

//////////////////////
func GetVideoList(sessionId, accessToken string, vResult *PlaybackVideoListResult) {
	if !CheckAuthValid(sessionId, accessToken, vResult.ApiResult) {
		return
	}

	videoList := make([]model.LiveVideo, 0)
	qErr := model.GetLiveVideoList(&videoList)
	if qErr != nil {
		vResult.SetCode(API_SERVER_ERROR)
		return
	}

	playbackVideoList := make([]PlaybackVideo, 0)
	for _, video := range videoList {
		playbackVideo := PlaybackVideo{
			User:       video.User.Name,
			Title:      video.Title,
			PublishId:  video.PublishId,
			CreateTime: video.CreateTime.Unix(),
		}

		playbackVideoList = append(playbackVideoList, playbackVideo)
	}

	vResult.VideoList = playbackVideoList
	vResult.SetOk()

	return
}

func GetPublishingList(sessionId, accessToken string, vResult *PublishingVideoListResult) {
	if !CheckAuthValid(sessionId, accessToken, vResult.ApiResult) {
		return
	}

	videoList := make([]model.LiveVideo, 0)
	qErr := model.GetLiveStreamList(&videoList)
	if qErr != nil {
		log.Error("get live video list error,", qErr.Error())
		vResult.SetCode(API_SERVER_ERROR)
		return
	}

	publishingVideoList := make([]PublishingVideo, 0)
	for _, video := range videoList {
		publishingVideo := PublishingVideo{
			User:       video.User.Name,
			Title:      video.Title,
			PublishId:  video.PublishId,
			CreateTime: video.CreateTime.Unix(),
		}

		publishingVideoList = append(publishingVideoList, publishingVideo)
	}

	vResult.VideoList = publishingVideoList
	vResult.SetOk()

	return
}

func GetMyVideoList(sessionId, accessToken string, vResult *PlaybackVideoListResult) {
	if !CheckAuthValid(sessionId, accessToken, vResult.ApiResult) {
		return
	}

	userId, gErr := model.GetSession(sessionId)
	if gErr != nil {
		log.Error("get session error,", gErr.Error())
		vResult.SetCode(API_SESSION_EXPIRED_ERROR)
		return
	}

	videoList := make([]model.LiveVideo, 0)
	qErr := model.GetMyLiveVideoList(userId, &videoList)
	if qErr != nil {
		vResult.SetCode(API_SERVER_ERROR)
		return
	}

	playbackVideoList := make([]PlaybackVideo, 0)
	for _, video := range videoList {
		playbackVideo := PlaybackVideo{
			User:       video.User.Name,
			Title:      video.Title,
			PublishId:  video.PublishId,
			CreateTime: video.CreateTime.Unix(),
		}

		playbackVideoList = append(playbackVideoList, playbackVideo)
	}

	vResult.VideoList = playbackVideoList
	vResult.SetOk()

	return
}

func GetStreamPlayResult(sessionId, accessToken, publishId string, pResult *StreamPlayResult) {
	if !CheckAuthValid(sessionId, accessToken, pResult.ApiResult) {
		return
	}

	if publishId == "" {
		pResult.SetFormatCode(API_PARAM_ERROR, "publish id is empty")
		return
	}

	liveVideo := model.LiveVideo{}

	qErr := model.GetLiveVideoByPublishId(publishId, &liveVideo)
	if qErr != nil {
		pResult.SetCode(API_SERVER_ERROR)
		return
	}

	if liveVideo.Id == 0 {
		pResult.SetCode(API_NO_VIDEO_FOUND_ERROR)
		return
	}

	playUrls, gErr := pilis.GetLivePlayUrl(liveVideo.StreamId)
	if gErr != nil {
		pResult.SetCode(API_SERVER_ERROR)
		return
	}

	pResult.PlayUrls = playUrls
	pResult.Orientation=liveVideo.Orientation
	pResult.SetOk()
	return
}

func GetVideoPlayResult(sessionId, accessToken, publishId string, pResult *VideoPlayResult) {
	if !CheckAuthValid(sessionId, accessToken, pResult.ApiResult) {
		return
	}

	if publishId == "" {
		pResult.SetFormatCode(API_PARAM_ERROR, "publish id is empty")
		return
	}

	liveVideo := model.LiveVideo{}

	qErr := model.GetLiveVideoByPublishId(publishId, &liveVideo)
	if qErr != nil {
		pResult.SetCode(API_SERVER_ERROR)
		return
	}

	if liveVideo.Id == 0 {
		pResult.SetCode(API_NO_VIDEO_FOUND_ERROR)
		return
	}

	playUrls, gErr := pilis.GetPlaybackUrl(liveVideo.StreamId, liveVideo.StartTime, liveVideo.EndTime)
	if gErr != nil {
		pResult.SetCode(API_SERVER_ERROR)
		return
	}

	pResult.PlayUrls = playUrls
	pResult.Orientation=liveVideo.Orientation
	pResult.SetOk()
	return
}
