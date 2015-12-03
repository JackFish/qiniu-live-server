package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//md5 hash to hex string
func Md5Hash(val string) string {
	h := md5.New()
	h.Write([]byte(val))
	output := h.Sum(nil)
	return hex.EncodeToString(output)
}

func UrlsafeBase64Encode(val string) string {
	return base64.URLEncoding.EncodeToString([]byte(val))
}

func UrlsafeBase64Decode(val string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(val)
}

//@param mobile
//@param pwd=md5(raw_pwd)
//@return md5(mobile:pwd:ts)
func CreateSessionId(mobile, pwd string) string {
	ts := time.Now().Unix()
	return Md5Hash(fmt.Sprintf("%s:%s:%d", mobile, pwd, ts))
}

//access_token=md5(auth_code+":"+ts+":"+auth_code)+":"+base64(ts)
//@param mobile
//@param authCode
//@param accessToken
//@return true|false
func IsAccessTokenValid(sessionId, accessToken string) (valid bool) {
	items := strings.Split(accessToken, ":")
	if len(items) != 2 {
		return
	}

	var ts string
	if val, bErr := UrlsafeBase64Decode(items[1]); bErr != nil {
		return
	} else {
		ts = string(val)

		if _, pErr := strconv.ParseInt(ts, 10, 64); pErr != nil {
			return
		} else {
			//check expire
			//expire := time.Unix(tsVal, 0).Add(time.Minute * 15)
			//if time.Now().After(expire) {
			//	return
			//}
		}
	}

	localHash := Md5Hash(fmt.Sprintf("%s:%s:%s", sessionId, ts, sessionId))
	valid = (items[0] == localHash)

	return
}
