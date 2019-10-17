package main

import (
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path"
	"time"
)

var logger *logrus.Entry

type ServerLogger struct {
	maxAge        time.Duration
	rotationTime  time.Duration
	logPath       string
	logFileName   string
	contextLogger *logrus.Logger
}

func (s *ServerLogger) configureLogToLocal() {
	s.contextLogger = logrus.New()
	baseLogPath := path.Join(s.logPath, s.logFileName)
	writer, err := rotatelogs.New(
		baseLogPath+".%Y%m%d",
		rotatelogs.WithLinkName(baseLogPath),        // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(s.maxAge),             // 文件最大保存时间
		rotatelogs.WithRotationTime(s.rotationTime), // 日志切割时间间隔
	)
	if err != nil {
		logrus.Fatalf("[视频转码服务]配置日志本地化操作失败, 错误原因: %v", err)
	}
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer,
		logrus.InfoLevel:  writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
	})
	s.contextLogger.AddHook(lfHook)
}

func NewLogger(maxAge, rotationTime time.Duration, logPath, logFileName string) *logrus.Entry {
	newLogger := &ServerLogger{
		maxAge:       maxAge,
		rotationTime: rotationTime,
		logPath:      logPath,
		logFileName:  logFileName,
	}
	newLogger.configureLogToLocal()
	newLogger.contextLogger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
	})
	logger = newLogger.contextLogger.WithFields(logrus.Fields{
		"serviceName": serviceName,
		"machineIP":   GetLocalIPAddress(),
		"version":     currentServiceVersion,
	})
	return logger
}
