package main

import (
	"net"
	"strings"
)

// 获取本机 IP 地址
func GetLocalIPAddress() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		logger.Fatal(err)
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	defer func() {
		if err := conn.Close(); err != nil {
			logger.Fatal(err)
		}
	}()
	localIPAddress := localAddr.IP
	return localIPAddress
}

// 字符串 format
func stringFormat(format string, args ...string) (formatString string) {
	return strings.NewReplacer(args...).Replace(format)
}

// 三元运算符
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
