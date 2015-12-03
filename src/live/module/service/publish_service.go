package service

import (
	"live/module/model"
	"strconv"
)

type StreamPublishStatus struct {
	ApiResult
}

func (this *StreamPublishStatus) SetOk() {
	this.Code = API_OK
	this.Desc = "stream is ready"
}

type StartPublishResult struct {
	ApiResult
	PublishId string `json:"publishId,omitempty"`
}

func (this *StartPublishResult) SetOk() {
	this.Code = API_OK
	this.Desc = "start publish success"
}

type StopPublishResult struct {
	ApiResult
}

func (this *StopPublishResult) SetOk() {
	this.Code = API_OK
	this.Desc = "stop publish success"
}

func StatusPublish(sessionId, accessToken, streamId string, pResult *StreamPublishStatus) {
	if _, valid := CheckAuthValid(sessionId, accessToken, &pResult.ApiResult); !valid {
		return
	}

	if streamId == "" {
		pResult.SetFormatCode(API_PARAM_ERROR, "stream id is empty")
		return
	}

	publishing, gErr := model.IsStreamPublishing(streamId)
	if gErr != nil {
		pResult.SetCode(API_SERVER_ERROR)
		return
	}

	if publishing {
		pResult.SetCode(API_STREAM_IS_TAKEN_ERROR)
		return
	}

	pResult.SetOk()
	return
}

func StartPublish(sessionId, accessToken, streamTitle, streamId, streamQualityStr, streamOrientationStr string, pResult *StartPublishResult) {
	if _, valid := CheckAuthValid(sessionId, accessToken, &pResult.ApiResult); !valid {
		return
	}

	if streamId == "" {
		pResult.SetFormatCode(API_PARAM_ERROR, "stream id is empty")
		return
	}

	if streamTitle == "" {
		pResult.SetFormatCode(API_PARAM_ERROR, "stream title is empty")
		return
	}

	streamQuality, pErr := strconv.ParseInt(streamQualityStr, 10, 64)
	if pErr != nil {
		pResult.SetFormatCode(API_PARAM_ERROR, "stream quality is invalid")
		return
	}

	streamOrientation, pErr := strconv.ParseInt(streamOrientationStr, 10, 64)
	if pErr != nil {
		pResult.SetFormatCode(API_PARAM_ERROR, "stream orientation is invalid")
		return
	}

	userId, gErr := model.GetSession(sessionId)
	if gErr != nil {
		pResult.SetCode(API_SERVER_ERROR)
		return
	}

	publishId, cErr := model.CreateNewPublish(userId, streamId, streamTitle, int(streamQuality), int(streamOrientation))
	if cErr != nil {
		pResult.SetCode(API_SERVER_ERROR)
		return
	}

	pResult.PublishId = publishId
	pResult.SetOk()

	return
}

func StopPublish(sessionId, accessToken, publishId string, pResult *StopPublishResult) {
	if _, valid := CheckAuthValid(sessionId, accessToken, &pResult.ApiResult); !valid {
		return
	}

	if publishId == "" {
		pResult.SetFormatCode(API_PARAM_ERROR, "publish id is empty")
		return
	}

	cErr := model.CompletePublish(publishId)
	if cErr != nil {
		pResult.SetCode(API_SERVER_ERROR)
		return
	}

	pResult.SetOk()

	return
}
