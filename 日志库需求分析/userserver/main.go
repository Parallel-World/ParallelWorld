package main

import (
	"studyGo/日志库/log"
)

// func test(){
// 	file := log.NewFileLog("D:/golang/Go/src/studyGo/日志库/userserver/main.go")
// 	file.LogDebug("this is a debug log")
// 	file.LogWarn("this is a warn log")
// }
// func test() {
// 	console := log.NewConsoleLog("xxx")
// 	console.LogConsoleLogDebug("this is a console log")
// 	console.LogConsoleLogWarn("this is a warn log")
// }

func test() {
	log := log.NewFileLog("c:/a.log") //log := log.NewConsoleLog("xxx")
	log.LogDebug("this is a debug file")
	log.LogWarn("this is a warn log")
}

func main() {
	test()
}
