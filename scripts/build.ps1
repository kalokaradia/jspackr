$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o bin/jspackr-linux-x64 ./src/main
$env:GOOS="linux"; $env:GOARCH="arm64"; go build -o bin/jspackr-linux-arm64 ./src/main

$env:GOOS="darwin"; $env:GOARCH="amd64"; go build -o bin/jspackr-darwin-x64 ./src/main
$env:GOOS="darwin"; $env:GOARCH="arm64"; go build -o bin/jspackr-darwin-arm64 ./src/main

$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o bin/jspackr-win32-x64.exe ./src/main
$env:GOOS="windows"; $env:GOARCH="arm64"; go build -o bin/jspackr-win32-arm64.exe ./src/main