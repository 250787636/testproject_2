package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

type LogFormatter struct{}

func (m *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	var msg string
	//entry.Logger.SetReportCaller(true) //HasCaller()为true才会有调用信息

	// "唯一标识"
	id := "唯一标识"
	// 日期
	timestamp := entry.Time.Format("20060102150405")
	// 响应是否成功
	isSuccess := filterSuccess(entry.Level.String())
	// "登录账号"
	account := "登录账号"
	// "IPv4"
	serverIPv4 := "IPv4"
	// "IPv6"
	serverIPv6 := "IPv6"
	// "访IPv4"
	accountIpv4 := "访IPv4"
	// "系统名称"
	systemName := "系统名称"

	if entry.HasCaller() {
		// "日志类型"
		logsType := filterLogType(entry.Caller.Function)
		// 打印信息
		//c.ClientIP()
		message := entry.Message
		if logsType == "登录日志" {

			// "登录类型"
			loginType := "登录类型"

			msg = fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|\n",
				id, timestamp, account, serverIPv4, serverIPv6, accountIpv4,
				logsType, systemName, loginType,
				message, isSuccess)
		} else {
			// "操作模块"
			useModule := "操作模块"
			// "操作类型"
			useType := "操作类型"
			// "操作结果"
			useResult := "操作结果"
			//"操作内容
			useContent := "操作内容"

			msg = fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|\n",
				id, timestamp, account, serverIPv4, serverIPv6, accountIpv4,
				logsType, systemName, useModule, useType, useContent, useResult,
				message, isSuccess)
		}
	} else {
		msg = fmt.Sprintf("%s|%s| %s\n", timestamp, entry.Level, entry.Message)
	}
	b.WriteString(msg)
	return b.Bytes(), nil
}
func main() {
	logrus.SetFormatter(&LogFormatter{})
	f, _ := os.OpenFile("mylog/test.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	logrus.SetReportCaller(true)
	// 显示文件和代码行数
	logrus.SetOutput(io.MultiWriter(f, os.Stdout))
	CheckLoginIn()
	MyLoginIn()
	Normal()
	NormalType()
	//logrus.Info("dwadawdadawd")
	//logrus.Error("error")
	//logrus.Warn("WARNING")
	//log := logrus.New()
	//log.SetFormatter(&LogFormatter{})
	//log.SetReportCaller(true)
	//log.SetOutput(os.Stdout)
	//log.Info("info")
	//log.Debug("Debug")
	//log.Warning("warning")
	//log.Error("error")
}

// 根据log类型进行成功或者失败的判断
func filterSuccess(str string) string {
	if str == "info" || str == "debug" {
		return "成功"
	} else if str == "warning" || str == "error" {
		return "失败"
	}
	return ""
}

// 根据方法名判断是登录日志还是操作日志
func filterLogType(str string) string {
	isLog := strings.Contains(strings.ToUpper(str), "LOG")
	if isLog {
		return "登录日志"
	}
	return "操作日志"
}

func CheckLoginIn() {
	logrus.Println("这是CheckLoginIn方法")
}

func MyLoginIn() {
	logrus.Println("这是MyLoginIn方法")
}
func Normal() {
	logrus.Println("这是Normal方法")
}
func NormalType() {
	logrus.Println("这是NormalType方法")
}

type MyLogs struct {
}
