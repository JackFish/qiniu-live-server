#Qiniu Live Server API 文档

##返回码列表

请求相关

|返回码|描述|值|
|------|-------------|-----------|
|API_OK|请求正常，且没有任何错误|1000|
|API_SERVER_ERROR|服务器内部处理错误，一般不会发生|1001|
|API_PARAM_ERROR|客户端请求参数错误，一般不会发生|1002|
|API_UNAUTHORIZED|客户端请求未授权，一般不会发生|1003|

回复相关

|返回码|描述|值|
|------|-------------|-----------|
|API_USER_NOT_FOUND|登陆时，用户不存在错误|1200|
|API_USER_EXISTS|注册时，用户已存储错误|1201|
|API_USER_PWD_ERROR|登陆时，用户密码错误|1202|
|API_STREAM_IS_TAKEN_ERROR|待推流时，流已经在推流|1301|
|API_NO_VIDEO_FOUND_ERROR|该视频没有找到|1401|

##请求回复说明

所有的API，如果请求执行失败，会返回如下格式的回复内容：

```
{
  "code": <ErrCode>,
  "desc": <ErrMsg>
}
```

例如：


```
{
  "code": 1002,
  "desc": "param error, mobile is empty"
}
```

```
{
  "code": 1003,
  "desc": "access not allowed"
}
```

```
{
  "code": 1200,
  "desc": "user not found"
}
```


如果请求成功，那么除返回成功的编码和描述信息外，还会返回额外API回复信息。


##注册

```
/signup?mobile=<Mobile>&pwd=<Password>&name=<UserName>&email=<Email>
```

返回码

```
API_OK
API_PARAM_ERROR
API_USER_EXISTS
API_SERVER_ERROR
```

请求成功返回值

```
{
  "code": 1000,
  "desc": "signup success"
}
```

##登录

```
/login?mobile=<Mobile>&pwd=<Password>
```

返回码

```
API_OK
API_USER_NOT_FOUND
API_USER_PWD_ERROR
API_SERVER_ERROR
```

请求成功返回值

```
{
  "code": 1000,
  "desc": "login success",
  "userName": "jemy",
  "sessionId": "5d42abbd429ae4a5f52b97dc931a1116"
}
```

##请求流信息

```
/get/stream?sessionId=<SessionId>&accessToken=<AccessToken>
```

返回码

```
API_OK
API_PARAM_ERROR
API_UNAUTHORIZED
API_SERVER_ERROR
```

请求成功返回值

```
{
  "code": 1000,
  "desc": "get stream success",
  "streamId": "z1.jinxinxin.563594c7eb6f92391700007a",
  "stream": "{\"id\":\"z1.jinxinxin.563594c7eb6f92391700007a\",\"createdAt\":\"2015-11-01T12:27:51.044+08:00\",\"updatedAt\":\"2015-11-01T12:27:51.044+08:00\",\"title\":\"563594c7eb6f92391700007a\",\"hub\":\"jinxinxin\",\"disabled\":false,\"publishKey\":\"6ac3e8a6a511816f\",\"publishSecurity\":\"dynamic\",\"hosts\":{\"publish\":{\"rtmp\":\"pili-publish.live.golanghome.com\"},\"live\":{\"hdl\":\"pili-live-hdl.live.golanghome.com\",\"hls\":\"pili-live-hls.live.golanghome.com\",\"http\":\"pili-live-hls.live.golanghome.com\",\"rtmp\":\"pili-live-rtmp.live.golanghome.com\"},\"playback\":{\"hls\":\"pili-playback.live.golanghome.com\",\"http\":\"pili-playback.live.golanghome.com\"}}}"
}
```

##查询流状态

```
/status/stream?sessionId=<SessionId>&accessToken=<AccessToken>&streamId=<StreamId>
```

返回码

```
API_OK
API_PARAM_ERROR
API_UNAUTHORIZED
API_SERVER_ERROR
API_STREAM_IS_TAKEN_ERROR
```

请求成功返回值

```
{
  "code": 1000,
  "desc": "stream is ready"
}
```


##开始推流

```
/start/publish?sessionId=<SessionId>&accessToken=<AccessToken>&streamId=<StreamId>&streamTitle=<streamTitle>
```

返回码

```
API_OK
API_PARAM_ERROR
API_UNAUTHORIZED
API_SERVER_ERROR
```

请求成功返回值

```
{
  "code": 1000,
  "desc": "start publish success",
  "publishId": "caf48ee98b0a2f739cef338d231cbe01"
}
```

##结束推流

```
/stop/publish?sessionId=<SessionId>&accessToken=<AccessToken>&publishId=<PublishId>
```

返回码

```
API_OK
API_PARAM_ERROR
API_UNAUTHORIZED
API_SERVER_ERROR
```

请求成功返回值

```
{
  "code": 1000,
  "desc": "stop publish success"
}
```

##获取所有正在推流的列表

```
/live/stream/list?sessionId=<SessionId>&accessToken=<AccessToken>
```

返回码

```
API_OK
API_PARAM_ERROR
API_UNAUTHORIZED
API_SERVER_ERROR
```

请求成功返回值

```
{
    "code": 1000,
    "desc": "get publishing list success",
    "videoList": [
        {
            "user": "jemygraw",
            "title": "cast demo 1",
            "publishId": "e945d80f3d48a983c4435d66df57d682",
            "createTime": 1446982171
        }
    ]
}
```


##获取所有推流完成的回放列表

```
/live/video/list?sessionId=<SessionId>&accessToken=<AccessToken>
```


返回码

```
API_OK
API_PARAM_ERROR
API_UNAUTHORIZED
API_SERVER_ERROR
```

请求成功返回值

```
{
    "code": 1000,
    "desc": "get video list success",
    "videoList": [
        {
            "user": "jemygraw",
            "title": "cast demo landscape",
            "publishId": "ba3e96a327637fa2b876f2f2fd9f3790",
            "createTime": 1446981713
        },
        {
            "user": "jemygraw",
            "title": "cast demo portrait",
            "publishId": "846f2b826cc4a92afdb9b5c5546aabc6",
            "createTime": 1446981607
        }
    ]
}
```

##获取我的推流完成的回放列表

```
/my/video/list?sessionId=<SessionId>&accessToken=<AccessToken>
```


返回码

```
API_OK
API_PARAM_ERROR
API_UNAUTHORIZED
API_SERVER_ERROR
```

请求成功返回值

```
{
    "code": 1000,
    "desc": "get video list success",
    "videoList": [
        {
            "user": "jemygraw",
            "title": "cast demo landscape",
            "publishId": "ba3e96a327637fa2b876f2f2fd9f3790",
            "createTime": 1446981713
        },
        {
            "user": "jemygraw",
            "title": "cast demo portrait",
            "publishId": "846f2b826cc4a92afdb9b5c5546aabc6",
            "createTime": 1446981607
        }
    ]
}
```

##获取选中流的播放地址

```
/get/play/stream?sessionId=<SessionId>&accessToken=<AccessToken>&publishId=<PublishId>
```
返回码

```
API_OK
API_PARAM_ERROR
API_UNAUTHORIZED
API_SERVER_ERROR
```

请求成功返回值

```
{
    "code": 1000,
    "desc": "get stream play url success",
    "orientation": 0,
    "playUrls": {
        "ORIGIN": "rtmp://pili-live-rtmp.live.golanghome.com/jinxinxin/563dea88d409d26e840013ce"
    }
}
```


##获取选中回放的播放地址


```
/get/play/video?sessionId=<SessionId>&accessToken=<AccessToken>&publishId=<PublishId>
```
返回码

```
API_OK
API_PARAM_ERROR
API_UNAUTHORIZED
API_SERVER_ERROR
```

请求成功返回值

```
{
    "code": 1000,
    "desc": "get video play url success",
    "orientation": 0,
    "playUrls": {
        "ORIGIN": "http://pili-playback.live.golanghome.com/jinxinxin/563dea88d409d26e840013ce.m3u8?start=1446981712&end=1446981731"
    }
}
```

##备注

**1. AccessToken的生成方法说明**

>AccessToken是直播业务服务端API验证请求合法性的字符串，由客户端根据一定的算法生成。由于本项目是演示项目，所以采用的API合法性验证方式可能比较简陋，实际开发
>过程中，请根据自身业务特点选用合适API验证方式。

本直播业务服务器端的AccessToken生成规则如下：

```
public static String md5Hash(String s) {
    char hexDigits[] = {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'};
    try {
        byte[] btInput = s.getBytes();
        MessageDigest mdInst = MessageDigest.getInstance("MD5");
        mdInst.update(btInput);
        byte[] md = mdInst.digest();
        int j = md.length;
        char str[] = new char[j * 2];
        int k = 0;
        for (int i = 0; i < j; i++) {
            byte byte0 = md[i];
            str[k++] = hexDigits[byte0 >>> 4 & 0xf];
            str[k++] = hexDigits[byte0 & 0xf];
        }
        return new String(str);
    } catch (Exception e) {
        e.printStackTrace();
        return null;
    }
}
    
public static String createAccessToken(String sessionId) {
    long timestamp = System.currentTimeMillis() / 1000;
    String ts = String.format("%d", timestamp);
    String encodedTs = Base64.encodeToString(ts.getBytes(), Base64.URL_SAFE);
    String toSign = String.format("%s:%s:%s", sessionId, ts, sessionId);
    String accessToken = String.format("%s:%s", Tools.md5Hash(toSign), encodedTs);
    return accessToken;
}
```