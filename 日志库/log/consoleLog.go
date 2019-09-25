package log

import "fmt"

type ConsoleLog struct {
}

func NewConsoleLog(file string) LogInterface {
	return &ConsoleLog{}
}

func (f *ConsoleLog) LogDebug(msg string) {
	fmt.Println("ConsoleLog", msg)
}

func (f *ConsoleLog) LogWarn(msg string) {
	fmt.Println("ConsoleLog", msg)
}
