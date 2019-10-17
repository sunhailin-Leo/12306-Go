#!/bin/bash
echo ">>>>>>>>>>>>>>>>>>>>>>>>> 正在打 Linux 环境包 >>>>>>>>>>>>>>>>>>>>>>>>>"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pkg/linux/go_12306 .
cp 12306_station.json pkg/linux/
echo ">>>>>>>>>>>>>>>>>>>>>>>>> Linux 环境包打包完成 >>>>>>>>>>>>>>>>>>>>>>>>>"

echo ">>>>>>>>>>>>>>>>>>>>>>>>> 正在打 Windows 环境包 >>>>>>>>>>>>>>>>>>>>>>>>>"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o pkg/win64/go_12306.exe .
cp 12306_station.json pkg/win64/
echo ">>>>>>>>>>>>>>>>>>>>>>>>> Windows 环境包打包完成 >>>>>>>>>>>>>>>>>>>>>>>>>"

echo ">>>>>>>>>>>>>>>>>>>>>>>>> 正在打 Mac 环境包 >>>>>>>>>>>>>>>>>>>>>>>>>"
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o pkg/darwin/go_12306 .
cp 12306_station.json pkg/darwin/
echo ">>>>>>>>>>>>>>>>>>>>>>>>> Mac 环境包打包完成 >>>>>>>>>>>>>>>>>>>>>>>>>"