package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"time"
)

func Md5Hash(val string) string {
	h := md5.New()
	h.Write([]byte(val))
	output := h.Sum(nil)
	return hex.EncodeToString(output)
}

func UrlsafeBase64Encode(val string) string {
	return base64.URLEncoding.EncodeToString([]byte(val))
}

func main() {
	sessionId := "ecea56e3931dbf0f294f0a09acd7908e"
	ts := fmt.Sprintf("%d", time.Now().Unix())

	localHash := Md5Hash(fmt.Sprintf("%s:%s:%s", sessionId, ts, sessionId))
	accessToken := fmt.Sprintf("%s:%s", localHash, UrlsafeBase64Encode(ts))

	fmt.Println(accessToken)
}
