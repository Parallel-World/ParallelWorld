package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	var c chan int
// 	fmt.Printf("c=%v\n", c)
// 	c = make(chan int, 10)
// 	fmt.Printf("c=%v\n", c)
// 	c <- 100
// 	data := <-c
// 	fmt.Println("data:", data)
// }

// func produce(c chan int) {
// 	c <- 1000
// 	fmt.Println("produce finished")
// }
// func consume(c chan int) {
// 	data := <-c
// 	fmt.Println("data:", data)
// }
// func main() {
// 	var c chan int
// 	fmt.Printf("c=%v\n", c)
// 	c = make(chan int)
// 	go produce(c)
// 	go consume(c)
// 	time.Sleep(time.Second * 3)
// }

// func hello(c chan bool){
// 	fmt.Println("hello goroutine")
// 	c <- true
// }
// func main(){
// 	var exitChan chan bool
// 	exitChan = make(chan bool)
// 	go hello(exitChan)
// 	fmt.Println("main thread terminate")
// 	<-exitChan
// }

// func sendData(sendch chan<- int) {
// 	sendch <- 10
// }
// func readData(sendch <-chan int) {
// 	data := <-sendch
// 	fmt.Println("data:", data)
// }
// func main() {
// 	chnl := make(chan int)
// 	go sendData(chnl)
// 	readData(chnl)
// }

// func producer(chnl chan int) {
// 	for i := 0; i < 10; i++ {
// 		chnl <- i
// 	}
// 	close(chnl)
// }
// func main() {
// 	ch := make(chan int)
// 	go producer(ch)
// 	for {
// 		v, ok := <-ch
// 		if ok == false {
// 			fmt.Println("chan is closed")
// 			break
// 		}
// 		fmt.Println("Received", v, ok)
// 	}
// }

// func produer(chnl chan int) {
// 	for i := 0; i < 10; i++ {
// 		chnl <- i
// 	}
// 	close(chnl)
// }
// func main() {
// 	ch := make(chan int)
// 	go produer(ch)
// 	for v := range ch {
// 		fmt.Println("receive:", v)
// 	}
// }

// func main() {
// 	ch := make(chan string, 2)
// 	ch <- "hello"
// 	ch <- "world"
// 	s1 := <-ch
// 	s2 := <-ch
// 	fmt.Println(s1, s2)
// }

// func write(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Println("successfully wrote", i, "to ch")
// 	}
// 	close(ch)
// }
// func main() {
// 	ch := make(chan int, 2)
// 	go write(ch)
// 	time.Sleep(2 * time.Second)
// 	for v := range ch {
// 		fmt.Println("read value", v, "from ch")
// 		time.Sleep(2 * time.Second)
// 	}
// }

// func main() {
// 	ch := make(chan string, 3)
// 	ch <- "naveen"
// 	ch <- "paul"
// 	fmt.Println("capacity is", cap(ch))
// 	fmt.Println("length is", len(ch))
// 	fmt.Println("read value", <-ch)
// 	fmt.Println("new length is", len(ch))
// }

// func process(i int, ch chan bool) {
// 	fmt.Println("started Goroutine ", i)
// 	time.Sleep(2 * time.Second)
// 	fmt.Printf("Goroutine %d ended\n", i)
// 	ch <- true
// }
// func main() {
// 	no := 3
// 	exitChan := make(chan bool, no)
// 	for i := 0; i < no; i++ {
// 		go process(i, exitChan)
// 	}
// 	for i := 0; i < no; i++ {
// 		<-exitChan
// 	}
// 	fmt.Println("All go routines finished executing")
// }

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Gorutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}
func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished executing")
}
