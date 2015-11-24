export GOPATH=$GOPATH:/Users/jemy/QiniuCloud/Projects/qiniu-live-server
gox -output="qlived_linux_amd64" -os="linux" -arch="amd64"
gox -output="qlived_linux_386" -os="linux" -arch="386"
gox -output="qlived_windows_amd64.exe" -os="windows" -arch="amd64"
gox -output="qlived_windows_386.exe" -os="windows" -arch="386"
gox -output="qlived_darwin_amd64" -os="darwin" -arch="amd64"
gox -output="qlived_darwin_386" -os="darwin" -arch="386"
