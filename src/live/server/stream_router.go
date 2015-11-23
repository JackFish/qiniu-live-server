package server

import (
	"live/module/service"
	"net/http"
)

//get stream info
//@param sessionId
//@param accessToken
func (this *LiveServer) serveGetStream(w http.ResponseWriter, req *http.Request) {
	gResult := service.GetStreamResult{}
	req.ParseForm()

	sessionId := req.FormValue("sessionId")
	accessToken := req.FormValue("accessToken")

	service.GetStream(sessionId, accessToken, &gResult)
	this.serveResultJson(w, req, http.StatusOK, &gResult)
}

//get stream status
//@param sessionId
//@param accessToken
//@param streamId
func (this *LiveServer) serveStatusStream(w http.ResponseWriter, req *http.Request) {
	sResult := service.StreamPublishStatus{}
	req.ParseForm()

	sessionId := req.FormValue("sessionId")
	accessToken := req.FormValue("accessToken")
	streamId := req.FormValue("streamId")

	service.StatusPublish(sessionId, accessToken, streamId, &sResult)
	this.serveResultJson(w, req, http.StatusOK, &sResult)
}

//start publish
//@param sessionId
//@param accessToken
//@param streamId
//@param streamTitle
func (this *LiveServer) serveStartPublish(w http.ResponseWriter, req *http.Request) {
	pResult := service.StartPublishResult{}
	req.ParseForm()

	sessionId := req.FormValue("sessionId")
	accessToken := req.FormValue("accessToken")
	streamTitle := req.FormValue("streamTitle")
	streamId := req.FormValue("streamId")
	streamQualityStr := req.FormValue("streamQuality")
	streamOrientationStr := req.FormValue("streamOrientation")

	service.StartPublish(sessionId, accessToken, streamTitle, streamId,
		streamQualityStr, streamOrientationStr, &pResult)
	this.serveResultJson(w, req, http.StatusOK, &pResult)
}

//stop publish
//@param sessionId
//@param accessToken
//@param publishId
func (this *LiveServer) serveStopPublish(w http.ResponseWriter, req *http.Request) {
	pResult := service.StopPublishResult{}
	req.ParseForm()

	sessionId := req.FormValue("sessionId")
	accessToken := req.FormValue("accessToken")
	publishId := req.FormValue("publishId")

	service.StopPublish(sessionId, accessToken, publishId, &pResult)
	this.serveResultJson(w, req, http.StatusOK, &pResult)
}
