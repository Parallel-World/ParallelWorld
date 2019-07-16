package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// var x int
// var wg sync.WaitGroup
// var mutex sync.Mutex

// func add() {
// 	for i := 0; i < 500; i++ {
// 		mutex.Lock()
// 		x = x + 1
// 		mutex.Unlock()
// 	}
// 	wg.Done()
// }
// func main() {
// 	wg.Add(2)
// 	go add()
// 	go add()
// 	wg.Wait()
// 	fmt.Println("x:", x)
// }

// var rwlock sync.RWMutex
// var wg sync.WaitGroup
// var x int

// func write() {
// 	rwlock.Lock()
// 	fmt.Println("write lock")
// 	x = x + 1
// 	time.Sleep(time.Second * 2)
// 	fmt.Println("write unlock")
// 	rwlock.Unlock()
// 	wg.Done()
// }
// func read(i int) {
// 	rwlock.RLock()
// 	fmt.Printf("goroutine:%d x=%d\n", i, x)
// 	time.Sleep(time.Second)
// 	rwlock.RUnlock()
// 	wg.Done()
// }
// func main() {
// 	wg.Add(1)
// 	go write()
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go read(i)
// 	}
// 	wg.Wait()
// }

// var rwlock sync.RWMutex
// var wg sync.WaitGroup
// var mutex sync.Mutex
// var x int

// func write() {
// 	for i := 0; i < 100; i++ {
// 		mutex.Lock()
// 		x = x + 1
// 		time.Sleep(time.Millisecond * 10)
// 		mutex.Unlock()
// 	}
// 	wg.Done()
// }
// func read(i int) {
// 	for i := 0; i < 100; i++ {
// 		mutex.Lock()
// 		time.Sleep(time.Millisecond)
// 		mutex.Unlock()
// 	}
// 	wg.Done()
// }
// func main() {
// 	start := time.Now().UnixNano()
// 	wg.Add(1)
// 	go write()
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go read(i)
// 	}
// 	wg.Wait()
// 	end := time.Now().UnixNano()
// 	cost := (end - start) / 1000 / 1000
// 	fmt.Println("cost:", cost, "ms")
// }

var x int32
var wg sync.WaitGroup
var mutex sync.Mutex

func addMutex() {
	for i := 0; i < 500; i++ {
		mutex.Lock()
		x = x + 1
		mutex.Unlock()
	}
}
func add() {
	for i := 0; i < 500; i++ {
		atomic.AddInt32(&x, 1)
	}
	wg.Done()
}
func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go add()
		go addMutex()
	}
	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 / 1000
	fmt.Println("x:", x, "cost:", cost, "ms")
}
