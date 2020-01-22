SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
echo "Build to Linux x64"
go build -o out/linux/wizz-homepage-go .
echo "Build completed"
pause
exit
