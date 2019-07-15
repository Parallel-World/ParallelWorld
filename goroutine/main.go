package main

import (
	"fmt"
	"runtime"
	"time"
)

// func hello() {
// 	fmt.Println("hello goroutine")
// }
// func test() {
// 	go hello()
// 	fmt.Println("main thread terminate")
// 	time.Sleep(time.Second)
// }
// func hello(i int) { fmt.Println("hello goroutine", i) }
// func test() {
// 	for i := 0; i < 10; i++ {
// 		go hello(i)
// 	}
// }
// func main() {
// 	test()
// 	time.Sleep(time.Second)
// }
var i int

func calc() {
	for {
		i++
	}
}
func main() {
	cpu := runtime.NumCPU()
	fmt.Println("cpu", cpu)
	runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		go calc()
	}
	time.Sleep(time.Hour)
}
