package server

import (
	"encoding/json"
	"fmt"
	"github.com/qiniu/log"
	"live/config"
	"net/http"
	"time"
)

type LiveServer struct {
	cfg      *config.ServerConfig
	handlers map[string]ServiceHandler
}

type ServiceHandler func(w http.ResponseWriter, req *http.Request)

func NewServer(cfg *config.ServerConfig) *LiveServer {
	serv := LiveServer{}
	serv.cfg = cfg
	serv.handlers = map[string]ServiceHandler{
		cfg.LocationPrefix + "/login":              serv.serveLogin,
		cfg.LocationPrefix + "/signup":             serv.serveSignup,
		cfg.LocationPrefix + "/get/stream":         serv.serveGetStream,
		cfg.LocationPrefix + "/status/stream":      serv.serveStatusStream,
		cfg.LocationPrefix + "/start/publish":      serv.serveStartPublish,
		cfg.LocationPrefix + "/stop/publish":       serv.serveStopPublish,
		cfg.LocationPrefix + "/live/stream/list":   serv.serveLivePublishingList,
		cfg.LocationPrefix + "/live/video/list":    serv.serveLiveVideoList,
		cfg.LocationPrefix + "/my/live/video/list": serv.serveMyVideoList,
		cfg.LocationPrefix + "/my/live/play/urls":  serv.serveMyLivePlayUrls,
		cfg.LocationPrefix + "/get/play/stream":    serv.serveGetLivePlayStream,
		cfg.LocationPrefix + "/get/play/video":     serv.serveGetLivePlayVideo,
	}
	return &serv
}

func (this *LiveServer) Listen() (err error) {
	log.Info("start server listening")
	endPoint := fmt.Sprintf("%s:%d", this.cfg.ListenHost, this.cfg.ListenPort)
	server := &http.Server{
		Addr:           endPoint,
		ReadTimeout:    time.Duration(this.cfg.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(this.cfg.WriteTimeout) * time.Second,
		MaxHeaderBytes: this.cfg.MaxHeaderBytes,
		Handler:        this,
	}
	listenErr := server.ListenAndServe()
	if listenErr != nil {
		err = listenErr
		return
	}
	return
}

func (this *LiveServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if handler, ok := this.handlers[path]; ok {
		handler(w, req)
	} else {
		this.serveNotFound(w, req)
	}
}

func (this *LiveServer) serveNotFound(w http.ResponseWriter, req *http.Request) {
	log.Error(req.Method, req.RequestURI, http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"error":"serivce not found"}`))
}

func (this *LiveServer) serveResultJson(w http.ResponseWriter, req *http.Request, code int, obj interface{}) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(code)
	jsonBytes, _ := json.Marshal(obj)
	w.Write([]byte(jsonBytes))
}

func (this *LiveServer) serveResultOctect(w http.ResponseWriter, req *http.Request, code int, contentType string, respData []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(code)
	w.Write([]byte(respData))
}
