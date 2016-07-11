package pilis

import (
	"errors"
	"fmt"
	"github.com/pili-engineering/pili-sdk-go/pili"
	"live/config"
)


//get stream
//@param stream id
//@return stream, code,err
//if code==404, try to create a new stream
func GetStream(streamId string) (streamData string, code int, err error) {
	crendentials := pili.NewCredentials(config.App.AccessKey, config.App.SecretKey)
	hub := pili.NewHub(crendentials, config.App.LiveHub)
	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		if v, ok := gErr.(*pili.ErrorInfo); ok {
			err = fmt.Errorf("get live stream error, %s", v.Message)
			code = v.ErrCode
		}else {
			err = fmt.Errorf("get live stream error, %s", gErr.Error())
		}
		return
	}

	streamJson, tErr := stream.ToJSONString()
	if tErr != nil {
		err = fmt.Errorf("get live stream error, parse error %s", tErr.Error())
		return
	}

	streamData = streamJson
	return
}

//create stream
//@return streamId, stream, err
func CreateDynamicStream() (streamId string, streamData string, err error) {
	crendentials := pili.NewCredentials(config.App.AccessKey, config.App.SecretKey)
	hub := pili.NewHub(crendentials, config.App.LiveHub)
	options := pili.OptionalArguments{
		PublishSecurity: "dynamic",
	}
	stream, cErr := hub.CreateStream(options)
	if cErr != nil {
		err = errors.New(fmt.Sprintf("create live stream error, %s", cErr.Error()))
		return
	}

	streamJson, tErr := stream.ToJSONString()
	if tErr != nil {
		err = errors.New("marshal live stream error")
		return
	}

	streamData = streamJson
	streamId = stream.Id

	return
}

func GetLivePlayUrl(streamId string) (livePlayUrls map[string]string, err error) {
	crendentials := pili.NewCredentials(config.App.AccessKey, config.App.SecretKey)
	hub := pili.NewHub(crendentials, config.App.LiveHub)
	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		err = errors.New(fmt.Sprintf("get live stream error, %s", gErr.Error()))
		return
	}
	livePlayUrls, err = stream.RtmpLiveUrls()
	return
}

//@param streamId
//@return rtmp, hls, flv live play urls
func GetLivePlayUrls(streamId string) (livePlayUrls map[string]string, err error) {
	crendentials := pili.NewCredentials(config.App.AccessKey, config.App.SecretKey)
	hub := pili.NewHub(crendentials, config.App.LiveHub)
	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		err = errors.New(fmt.Sprintf("get live stream error, %s", gErr.Error()))
		return
	}

	livePlayUrls = make(map[string]string)

	rtmpLivePlayUrls, gErr := stream.RtmpLiveUrls()
	if gErr != nil {
		err = errors.New(fmt.Sprintf("get live stream rtmp play url error, %s", gErr.Error()))
		return
	}

	hlsLivePlayUrls, gErr := stream.HlsLiveUrls()
	if gErr != nil {
		err = errors.New(fmt.Sprintf("get live stream hls play url error, %s", gErr.Error()))
		return
	}

	flvLivePlayUrls, gErr := stream.HttpFlvLiveUrls()
	if gErr != nil {
		err = errors.New(fmt.Sprintf("get live stream flv play url error, %s", gErr.Error()))
		return
	}

	if v, ok := rtmpLivePlayUrls["ORIGIN"]; ok {
		livePlayUrls["RTMP"] = v
	}

	if v, ok := hlsLivePlayUrls["ORIGIN"]; ok {
		livePlayUrls["HLS"] = v
	}

	if v, ok := flvLivePlayUrls["ORIGIN"]; ok {
		livePlayUrls["FLV"] = v
	}

	return
}

func GetPlaybackUrl(streamId string, startTime, endTime int64) (playbackUrls map[string]string, err error) {
	crendentials := pili.NewCredentials(config.App.AccessKey, config.App.SecretKey)
	hub := pili.NewHub(crendentials, config.App.LiveHub)
	stream, gErr := hub.GetStream(streamId)
	if gErr != nil {
		err = errors.New(fmt.Sprintf("get live stream error, %s", gErr.Error()))
		return
	}

	playbackUrls, err = stream.HlsPlaybackUrls(startTime, endTime)
	return
}
