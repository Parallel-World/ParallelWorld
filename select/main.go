package main

import (
	"fmt"
	"time"
)

// func server1(ch chan string) {
// 	time.Sleep(time.Second * 3)
// 	ch <- "response from server1"
// }
// func server2(ch chan string) {
// 	time.Sleep(time.Second)
// 	ch <- "response from server2"
// }
// func main() {
// 	output1 := make(chan string)
// 	output2 := make(chan string)
// 	go server1(output1)
// 	go server2(output2)
// 	// s1 := <-output1
// 	// fmt.Println("s1:", s1)
// 	// s2 := <-output2
// 	// fmt.Println("s2:", s2)
// 	select {
// 	case s1 := <-output1:
// 		fmt.Println("s1:", s1)
// 	case s2 := <-output2:
// 		fmt.Println("s2:", s2)
// 	}
// }

func write(ch chan string) {
	for {
		select {
		case ch <- "hello":
			fmt.Println("write succ")
		default:
			fmt.Println("channel is full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
func main() {
	output1 := make(chan string, 10)
	go write(output1)
	for s := range output1 {
		fmt.Println("recv:", s)
		time.Sleep(time.Second)
	}
}
