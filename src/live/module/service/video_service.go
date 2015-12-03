package service

import (
	"github.com/qiniu/log"
	"live/module/model"
	"live/module/service/pilis"
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
	VideoList []PublishingVideo `json:"videoList,omitempty"`
}

func (this *PublishingVideoListResult) SetOk() {
	this.Code = API_OK
	this.Desc = "get publishing list success"
}

//////////////////////

type VideoPlayResult struct {
	ApiResult
	Orientation int               `json:"orientation"`
	PlayUrls    map[string]string `json:"playUrls,omitempty"`
}

func (this *VideoPlayResult) SetOk() {
	this.Code = API_OK
	this.Desc = "get video play url success"
}

//////////////////////

type StreamPlayResult struct {
	ApiResult
	Orientation int               `json:"orientation,omitempty"`
	PlayUrls    map[string]string `json:"playUrls,omitempty"`
}

func (this *StreamPlayResult) SetOk() {
	this.Code = API_OK
	this.Desc = "get stream play url success"
}

///////////////////////

type MyLivePlayUrlsResult struct {
	ApiResult
	LivePlayUrls map[string]string `json:"livePlayUrls,omitempty"`
}

func (this *MyLivePlayUrlsResult) SetOk() {
	this.Code = API_OK
	this.Desc = "get stream live play urls success"
}

//////////////////////
func GetVideoList(sessionId, accessToken string, vResult *PlaybackVideoListResult) {
	if _, valid := CheckAuthValid(sessionId, accessToken, &vResult.ApiResult); !valid {
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
	if _, valid := CheckAuthValid(sessionId, accessToken, &vResult.ApiResult); !valid {
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
	userId, valid := CheckAuthValid(sessionId, accessToken, &vResult.ApiResult)
	if !valid {
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
	if _, valid := CheckAuthValid(sessionId, accessToken, &pResult.ApiResult); !valid {
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
	pResult.Orientation = liveVideo.Orientation
	pResult.SetOk()
	return
}

func GetVideoPlayResult(sessionId, accessToken, publishId string, pResult *VideoPlayResult) {
	if _, valid := CheckAuthValid(sessionId, accessToken, &pResult.ApiResult); !valid {
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
	pResult.Orientation = liveVideo.Orientation
	pResult.SetOk()
	return
}

func GetMyLivePlayUrls(sessionId, accessToken string, pResult *MyLivePlayUrlsResult) {
	userId, valid := CheckAuthValid(sessionId, accessToken, &pResult.ApiResult)
	if !valid {
		return
	}

	streamId, gErr := model.GetStreamIdOfUser(userId)
	if gErr != nil {
		pResult.SetCode(API_SERVER_ERROR)
		return
	}

	playUrls, gErr := pilis.GetLivePlayUrls(streamId)
	if gErr != nil {
		pResult.SetCode(API_SERVER_ERROR)
		return
	}

	pResult.LivePlayUrls = playUrls
	pResult.SetOk()

	return
}
