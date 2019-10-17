package main

import (
	"os"
	"time"
)

const (
	serviceName           string = "12306 服务"
	currentServiceVersion string = "v0.1.0"
	logMaxAge                    = time.Hour * 24
	logRotationTime              = time.Hour * 24
	logPath               string = ""
	logFileName           string = "12306-train.log"
	stationJson           string = "12306_station.json"
	userAgent             string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"
)

func main() {
	// 初始化日志
	logger = NewLogger(logMaxAge, logRotationTime, logPath, logFileName)
	// 初始化
	loadJsonStation()

	// 命令行初始化
	commandLineInit()
	commands := Commands
	args := os.Args
	for _, cmd := range commands {
		if cmd.Run != nil && cmd.Name() == args[1] {
			err := cmd.Flag.Parse(args[2:])
			if err != nil {
				os.Exit(1)
			}
			args = cmd.Flag.Args()
			if len(args) > 0 {
				os.Exit(cmd.Run(args))
			}
			logger.Errorf("[12306查询助手]命令参数错误!")
			break
		}
	}
}
