package logger

import (
	"testing"
)

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "D:/golang/Go/src/studyGo/日志库/logger/", "test") //运行单元测试函数，测试Logger，级别，路径，名字
	logger.Debug("user id[%d] is come from china", 2321)
	logger.Warn("test warn log")
	logger.Fatal("test faral log")
	logger.Close()
}

func TestConsoleLogger(t *testing.T) {
	logger := NewConsoleLogger(LogLevelDebug) //控制台test
	logger.Debug("user id[%d] is come from china", 2321)
	logger.Warn("test warn log")
	logger.Fatal("test faral log")
	logger.Close()
}
