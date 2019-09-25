package main

import (
	"studyGo/日志库/logger"
	"time"
)

func initLogger(name string, logPath, LogName string, level string) (err error) {
	m := make(map[string]string, 8)
	m["log_path"] = logPath
	m["log_name"] = "user_server"
	m["log_level"] = level
	err = logger.InitLogger(name, m) //选择打印到控制台还是文件中
	if err != nil {
		return
	}
	logger.Debug("init logger success")
	return
}

func Run() {
	for {
		logger.Debug("user server is running")
		time.Sleep(time.Second)
	}
}

func main() {
	initLogger("file", "D:\\golang\\Go\\src\\studyGo\\日志库\\log\\", "user_server", "debug")
	Run()
	return
}
