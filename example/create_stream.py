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

stream=hub.create_stream(publishSecurity="dynamic")
print(stream.id)
print(stream.to_json())

rtmp_url=stream.rtmp_publish_url()
print(rtmp_url)


hls_urls=stream.hls_live_urls()
print(hls_urls)

