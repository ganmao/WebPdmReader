REM 交叉编译不支持CGO，所以禁用
SET CGO_ENABLED=0

SET GOOS=windows
SET GOARCH=amd64
go build -o WebPdmReader.exe

SET GOOS=linux
SET GOARCH=amd64
go build -o WebPdmReader.linux

SET GOOS=darwin
SET GOARCH=amd64
go build -o WebPdmReader.mac
