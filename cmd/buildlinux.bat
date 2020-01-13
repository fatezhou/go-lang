set GOARCH=amd64
set GOOS=linux
go build -o "../main" -ldflags="-s -w"