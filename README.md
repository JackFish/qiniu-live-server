#七牛直播demo服务端程序

##简介

该项目是七牛直播demo客户端 [QLive for Android](https://github.com/qiniudemo/qiniu-live-android) 和 [PLdemo for iOS](https://github.com/qiniudemo/qiniu-live-iOS) 的直播业务服务器端。

该项目是使用Golang编写的直播业务服务器端，可以直接运行，为了方便大家测试使用，可以使用已编译版本，[点击下载](http://devtools.qiniu.com/qiniu-live-server-v1.zip)。

该项目主要功能是提供客户端推流和观看直播视频的API接口，模拟了常规APP开发中所需要的业务服务器端的功能。

##使用方式

该项目需要MySQL的支持，所以首先系统得安装有MySQL，其他的都不需要。

创建数据库和表的SQL脚本为`database.sql`，解压下载后的zip包里面即有。

压缩包中提供了预编译好的支持`Linux`,`Windows`和`Mac`的可执行文件，根据自己的系统选择对应可执行文件。

压缩包中的`qlived.conf`是直播业务服务器端的配置文件，其中的内容安装格式填写即可，具体说明如下。

```
{
    "server":{
	    "listen_host": "0.0.0.0",
	    "listen_port": 9090,
	    "read_timeout": 60,
	    "write_timeout": 60,
	    "max_header_bytes": 4096
	},
	"app":{
		"live_hub": "hub-name",
		"access_key": "",
		"secret_key": "",
		"log_file":"run0.log",
		"log_level" : "debug"
	},
	"orm":{
		"driver_name":"mysql",
		"data_source":"root:root@tcp(localhost:3306)/qlive?charset=utf8&loc=Asia%2FShanghai",
		"max_idle_conn":30,
		"max_open_conn":50,
		"debug_mode":true
	}
}
```

**server节点**：

|字段|说明|默认|
|-----|------|-------|
|listen_host|监听地址，可以不用改|默认`0.0.0.0`|
|listen_port|监听端口，根据需要设置|默认`9090`|
|read_timeout|连接建立，读取数据超时时间，单位秒，可以不设置|默认`60`|
|write_timeout|连接建立，写入数据超时时间，单位秒，可以不设置|默认`60`|
|max_header_bytes|服务端允许的最大头部大小，可以不设置|默认4KB|

**app节点**

|字段|说明|默认|
|-----|------|--------|
|live_hub|直播服务的hub名称|根据注册的账号对应的hub名称填写|
|access_key|直播账号的`AccessKey`|根据注册的账号对应的AccessKey填写|
|secret_key|直播账号的`SecretKey`|根据注册的账号对应的SecretKey填写|
|log_file|日志文件名称，如果名称中有相对路径，那么该相对路径必须存在|指定为本地文件|
|log_level|日志级别，可以为`debug`,`error`,`info`等|默认为`debug`|


**orm节点**

该节点指定MySQL的相关信息，根据实际的MySQL配置修改用户名和密码即可，其他保持不变。

运行直播业务服务器端：

```
./qlived -c qlived.conf
```

##API说明

请看[API_DOC](API_DOC.md)
