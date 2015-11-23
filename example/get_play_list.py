from pili import *

access_key=""
secret_key=""

hub_name="jinxinxin"
rtmp_publish_host="pili-publish.live.golanghome.com"
hls_live_play_host="pili-live-hls.live.golanghome.com"

#publish address
#rtmp://pili-publish.live.golanghome.com/jinxinxin/test?key=abc

credentials=Credentials(access_key,secret_key)
hub=Hub(credentials,hub_name)

#z1.jinxinxin.56370fd1d409d29ff00004b4
#z1.jinxinxin.563594c7eb6f92391700007a
stream=hub.get_stream("z1.jinxinxin.563594c7eb6f92391700007a")
hls_urls=stream.hls_live_urls()
print(hls_urls)

rtmp_urls=stream.rtmp_live_urls()
print(rtmp_urls)

start=1446540177
end=1446540223
urls=stream.hls_playback_urls(start,end)
for k in urls:
	print(k,":",urls[k])



