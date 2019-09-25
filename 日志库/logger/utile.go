package logger

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

type LogData struct {
	Message      string
	TimeStr      string
	LevelStr     string
	Filename     string
	FuncName     string
	LineNo       int
	WarnAndFatal bool
}

//获取当前行号
func GetLineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4) //获取行号,数字表示栈的深度，0时是本身，1时是上层GetLineInfo,2时是
	//pc:程序指令计数器,计算运行到那个位置
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name() //通过pc获取到执行的函数.name获取函数名
		lineNo = line
	}
	return
}

/*
当业务调用打日志的方法时，把日志相关的数据写入到chan中
然后用一个后台的线程不断从chan里面获取这些日志，最终写到文件
*/

func writeLog(level int, format string, args ...interface{}) *LogData {
	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(level)
	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName) //只返回文件最后一个路径
	funcName = path.Base(funcName)
	msg := fmt.Sprintf(format, args...)

	logData := &LogData{
		Message:      msg,
		TimeStr:      nowStr,
		LevelStr:     levelStr,
		Filename:     fileName,
		LineNo:       lineNo,
		WarnAndFatal: false,
	}
	if level == LogLevelError || level == LogLevelWarn || level == LogLevelFatal {
		logData.WarnAndFatal = true
	}
	return logData
	// fmt.Fprintf(file, "%s %s (%s:%s %d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
}
