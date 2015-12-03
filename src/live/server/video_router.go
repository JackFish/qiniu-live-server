package server

import (
	"live/module/service"
	"net/http"
)

//get all live video list
//@param sessionId
//@param accessToken
func (this *LiveServer) serveLiveVideoList(w http.ResponseWriter, req *http.Request) {
	vResult := service.PlaybackVideoListResult{}

	req.ParseForm()

	sessionId := req.FormValue("sessionId")
	accessToken := req.FormValue("accessToken")

	service.GetVideoList(sessionId, accessToken, &vResult)
	this.serveResultJson(w, req, http.StatusOK, &vResult)
}

//get live publishing list
//@param sessionId
//@param accessToken
func (this *LiveServer) serveLivePublishingList(w http.ResponseWriter, req *http.Request) {
	vResult := service.PublishingVideoListResult{}

	req.ParseForm()
	sessionId := req.FormValue("sessionId")
	accessToken := req.FormValue("accessToken")

	service.GetPublishingList(sessionId, accessToken, &vResult)
	this.serveResultJson(w, req, http.StatusOK, &vResult)
}

//get all my video list
//@param sessionId
//@param accessToken
func (this *LiveServer) serveMyVideoList(w http.ResponseWriter, req *http.Request) {
	vResult := service.PlaybackVideoListResult{}

	req.ParseForm()

	sessionId := req.FormValue("sessionId")
	accessToken := req.FormValue("accessToken")

	service.GetMyVideoList(sessionId, accessToken, &vResult)
	this.serveResultJson(w, req, http.StatusOK, &vResult)
}

//get my live publish address
//@param sessionId
//@param accessToken
func (this *LiveServer) serveMyLivePlayUrls(w http.ResponseWriter, req *http.Request) {
	vResult := service.MyLivePlayUrlsResult{}

	req.ParseForm()

	sessionId := req.FormValue("sessionId")
	accessToken := req.FormValue("accessToken")

	service.GetMyLivePlayUrls(sessionId, accessToken, &vResult)
	this.serveResultJson(w, req, http.StatusOK, &vResult)
}

//get the play url for the publishing stream
//@param sessionId
//@param accessToken
//@param publishId
func (this *LiveServer) serveGetLivePlayStream(w http.ResponseWriter, req *http.Request) {
	pResult := service.StreamPlayResult{}

	req.ParseForm()

	sessionId := req.FormValue("sessionId")
	accessToken := req.FormValue("accessToken")
	publishId := req.FormValue("publishId")

	service.GetStreamPlayResult(sessionId, accessToken, publishId, &pResult)
	this.serveResultJson(w, req, http.StatusOK, &pResult)
}

//get the play url for the live video
//@param sessionId
//@param accessToken
//@param publishId
func (this *LiveServer) serveGetLivePlayVideo(w http.ResponseWriter, req *http.Request) {
	pResult := service.VideoPlayResult{}

	req.ParseForm()

	sessionId := req.FormValue("sessionId")
	accessToken := req.FormValue("accessToken")
	publishId := req.FormValue("publishId")

	service.GetVideoPlayResult(sessionId, accessToken, publishId, &pResult)
	this.serveResultJson(w, req, http.StatusOK, &pResult)
}
