| 加速域名                            | IN QINIU CNAME                         
|------------------------------------|----------------------------------------
| pili-publish.live.golanghome.com   | 437py3.publish.z1.pili.qiniudns.com
| pili-live-rtmp.live.golanghome.com | 437py3.live-rtmp.z1.pili.qiniudns.com
| pili-live-hdl.live.golanghome.com  | 437py3.live-hdl.z1.pili.qiniudns.com 
| pili-live-hls.live.golanghome.com  | 437py3.live-hls.z1.pili.qiniudns.com 
| pili-playback.live.golanghome.com  | 437py3.playback.z1.pili.qiniudns.com 
| pili-media.live.golanghome.com     | 437py3.media.z1.pili.qiniudns.com    
| pili-vod.live.golanghome.com       | 437py3.vod.z1.pili.qiniudns.com      
| pili-static.live.golanghome.com    | 437py3.static.z1.pili.qiniudns.com

该app模拟的是UGC场景下的用户推流

每次推流都使用新的流地址和dynamic的方式

static的推流地址

rtmp://{rtmp_publish_host}/{hub_name}/{stream_title}?key={public_key}



