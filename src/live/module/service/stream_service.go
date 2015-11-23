package service

import (
	"live/module/model"
	"live/module/service/pilis"
)

type GetStreamResult struct {
	ApiResult
	StreamId string `json:"streamId,omitempty"`
	Stream   string `json:"stream,omitempty"`
}

func (this *GetStreamResult) SetOk() {
	this.Code = API_OK
	this.Desc = "get stream success"
}

func GetStream(sessionId, accessToken string, gResult *GetStreamResult) {
	if !CheckAuthValid(sessionId, accessToken, gResult.ApiResult) {
		return
	}

	userId, gErr := model.GetSession(sessionId)
	if gErr != nil {
		gResult.SetCode(API_SERVER_ERROR)
		return
	}

	//get stream by user
	streamId, gErr := model.GetStreamIdOfUser(userId)
	if gErr != nil {
		gResult.SetCode(API_SERVER_ERROR)
		return
	}

	if streamId == "" {
		newStreamId, newStreamData, newStreamErr := pilis.CreateDynamicStream()
		if newStreamErr != nil {
			gResult.SetCode(API_SERVER_ERROR)
			return
		}
		//update user stream
		sErr := model.SetStreamIdOfUser(userId, newStreamId)
		if sErr != nil {
			gResult.SetCode(API_SERVER_ERROR)
			return
		}
		gResult.StreamId = streamId
		gResult.Stream = newStreamData
	} else {
		streamData, sErr := pilis.GetStream(streamId)
		if sErr != nil {
			gResult.SetCode(API_SERVER_ERROR)
			return
		}
		gResult.StreamId = streamId
		gResult.Stream = streamData
	}

	gResult.SetOk()
	return
}
