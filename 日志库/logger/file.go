package logger

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

//2018/3/26 0:02.383 DEBUG logDebug(源文件名).go:29(行号) this is a debug log
//2006-01-02 15:04:05.999
type FileLogger struct {
	level         int    //日志级别
	logPath       string //日志路径
	logName       string //日志名称
	file          *os.File
	warnFile      *os.File
	LogDataChan   chan *LogData
	logSplitType  int
	logSplitSize  int64
	lastSplitHour int
}

func NewFileLogger(config map[string]string) (log LogInterface, err error) {
	logPath, ok := config["log_path"]
	if !ok {
		err := fmt.Errorf("not found log_path")
		return nil, err
	}
	logName, ok := config["log_name"]
	if !ok {
		err := fmt.Errorf("not found log_name")
		return nil, err
	}
	logLevel, ok := config["log_path"]
	if !ok {
		err := fmt.Errorf("not found log_level")
		return nil, err
	}

	logChanSize, ok := config["log_chan_size"]
	if !ok {
		logChanSize = "50000"
	}

	var logSplitType int = LogSplitTypeHour
	var logSplitSize int64
	logSplitStr, ok := config["log_split_type"]
	if !ok {
		logSplitStr = "hour"
	} else {
		if logSplitStr == "size" {
			logSplitSizeStr, ok := config["log_split_size"]
			if !ok {
				logSplitSizeStr = "104857600" //100M
			}
			logSplitSize, err = strconv.ParseInt(logSplitSizeStr, 10, 64) //把logSplitSize转换成10进制64位数
			if err != nil {
				logSplitSize = 104857600
			}
			logSplitType = LogSplitTypeSize
		} else {
			logSplitType = LogSplitTypeHour
		}
	}

	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000
	}
	level := getLogLevel(logLevel)
	log = &FileLogger{
		level:         level,
		logPath:       logPath,
		logName:       logName,
		LogDataChan:   make(chan *LogData, chanSize),
		logSplitSize:  logSplitSize,
		logSplitType:  logSplitType,
		lastSplitHour: time.Now().Hour(),
	}
	log.Init()
	return
}

func (f *FileLogger) Init() {
	filename := fmt.Sprintf("%s%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed,err:%v", filename, err))
	}
	f.file = file

	//写错误日志和fatal日志的文件
	filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open faile %s failed,err:%v", filename, err))
	}
	f.warnFile = file
	go f.writeLogBackgroud()
}

func (f *FileLogger) checkSplitFile(warnFile bool) {
	if f.logSplitType == LogSplitTypeHour {
		now := time.Now()
		hour := now.Hour()
		if hour == f.lastSplitHour{
			return 
		}
		var backupFilename string
		if warnFile {
			backupFilename = fmt.Sprintf("%s/%s.log.wf_%04d%02d%02d%02d",
			f.logPath,f.logName,now.Year(),now.Month(),now.Day(),f.lastSplitHour
		}
	}
}

func (f *FileLogger) writeLogBackgroud() {
	for logData := range f.LogDataChan {
		var file *os.File = f.file
		if logData.WarnAndFatal {
			file = f.warnFile
		}
		f.checkSplitFile(file)
		fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n",
			logData.TimeStr, logData.LevelStr, logData.Filename,
			logData.FuncName, logData.LineNo, logData.Message)
	}
}

func (f *FileLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	logData := writeLog(LogLevelDebug, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	logData := writeLog(LogLevelTrace, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logData := writeLog(LogLevelInfo, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	logData := writeLog(LogLevelWarn, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	logData := writeLog(LogLevelError, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	logData := writeLog(LogLevelFatal, format, args...)
	select {
	case f.LogDataChan <- logData:
	default:
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
