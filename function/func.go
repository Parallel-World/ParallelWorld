package main

import (
	"fmt"
	"time"
)

// func sayHello() {
// 	fmt.Println("HelloWorld")
// }

// func add(a int, b int) int {
// 	sum := a + b
// 	return sum
// }

// func calc(a, b int) (int, int) {
// 	sum := a + b
// 	sub := a - b
// 	return sum, sub
// }

// func calc_v1(b ...int) int {
// 	sum := 0
// 	for i := 0; i < len(b); i++ {
// 		sum = sum + b[i]
// 	}
// 	return sum
// }

// func calc_v2(a int, b ...int) int {
// 	sum := a
// 	for i := 0; i < len(b); i++ {
// 		sum = sum + b[i]
// 	}
// 	return sum
// }

// func testDefer1() {
// 	defer fmt.Println("Hello")
// 	defer fmt.Println("Hello1")
// 	defer fmt.Println("Hello2")
// 	fmt.Println("Parallel")
// 	fmt.Println("World")
// }

// func testDefer2() {
// 	for i := 0; i < 5; i++ {
// 		defer fmt.Println(i)
// 	}
// }

// func testDefer3() {
// 	i := 0
// 	defer fmt.Println(i)
// 	i = 1000
// 	fmt.Println(i)
// }

// func add(a, b int) int {
// 	return a + b
// }

// func sub(a, b int) int {
// 	return a - b
// }

// func calc(a, b int, op func(int, int) int) int {
// 	return op(a, b)
// }

// func testDefer5() {
// 	sum := calc(100, 300, add)
// 	sub := calc(100, 300, sub)
// 	fmt.Printf("sum-%d sub-%d\n", sum, sub)
// }

// func Adder() func(int) int {
// 	var x int
// 	return func(d int) int {
// 		x += d
// 		return x
// 	}
// }
// func testClosure1() {
// 	f := Adder()
// 	ret := f(1)
// 	fmt.Printf("ret=%d\n", ret)
// 	ret = f(20)
// 	fmt.Printf("ret=%d\n", ret)
// 	ret = f(300)
// 	fmt.Printf("ret=%d\n", ret)
// }

// func add(base int) func(int) int {
// 	return func(i int) int {
// 		base += i
// 		return base
// 	}
// }
// func testClosure1() {
// 	tmp1 := add(10)
// 	fmt.Println(tmp1(1), tmp1(2))
// 	tmp2 := add(100)
// 	fmt.Println(tmp2(1), tmp2(2))
// }

// func test(suffix string) func(string) string {
// 	return func(name string) string {
// 		if !strings.HasSuffix(name, suffix) {
// 			return name + suffix
// 		}
// 		return name
// 	}
// }
// func testClosure1() {
// 	func1 := test(".bmp")
// 	func2 := test(".jpg")
// 	fmt.Println(func1("test"))
// 	fmt.Println(func2("test"))
// }

// func calc(base int) (func(int) int, func(int) int) {
// 	add := func(i int) int {
// 		base += i
// 		return base
// 	}
// 	sub := func(i int) int {
// 		base -= i
// 		return base
// 	}
// 	return add, sub
// }
// func testClosure1() {
// 	f1, f2 := calc(10)
// 	fmt.Println(f1(1), f2(2))
// 	fmt.Println(f1(3), f2(4))
// 	fmt.Println(f1(5), f2(6))
// 	fmt.Println(f1(7), f2(8))
// }

func testClosure1() {
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}

func main() {
	// sayHello()
	// s := add(100, 200)
	// fmt.Println(s)
	// sum, sub := calc(100, 80)
	// fmt.Println(sum, sub)
	// sum := calc_v1(10, 20)
	// fmt.Printf("sum=%d\n", sum)
	// sum := calc_v2(10, 10, 10)
	// fmt.Println(sum)
	// testDefer3()
	// input := bufio.NewScanner(os.Stdin)
	// input.Scan()
	// line := input.Text()
	// fmt.Println(line)
	// testDefer5()
	testClosure1()
}
