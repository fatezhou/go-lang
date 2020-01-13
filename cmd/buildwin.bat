set GOOS=windows

go build -o "../main.exe" -ldflags "-X 'main.BuildTime=%date:~0,10% %time%' -X main.Debug=false"
