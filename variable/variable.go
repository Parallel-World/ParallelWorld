package main

import (
	"fmt"
)

// func testvariable() {
// 	var a int32
// 	fmt.Printf("the addr of a:%p", &a)
// }

// func testvariable() {
// 	var a int32
// 	a = 156
// 	fmt.Printf("the addr of a:%p,a:%d\n", &a, a)
// 	var b *int32
// 	fmt.Printf("the addr of b:%p,b:%v\n", &b, b)
// 	if b==nil{
// 		fmt.Println("这个指针变量为空")
// 	}
// 	b = &a
// 	fmt.Printf("the addr of b:%p,b:%v\n", &b, b)
// }

// func testvariable() {
// 	var a int = 200
// 	var b *int = &a
// 	fmt.Printf("b指向的地址存储的值为:%d\n", *b)
// 	*b = 1000
// 	fmt.Printf("b指向的地址存储的值为:%d\n", *b)
// 	fmt.Printf("a=%d\n", a)
// }

// func modify(a *int) {
// 	*a = 100
// }
// func testvariable() {
// 	var b int = 10
// 	p := &b
// 	modify(p)
// 	fmt.Printf("b:%d\n", b)
// }

// func modify(a *[3]int) {
// 	(*a)[0] = 100
// }

// func testvariable() {
// 	var b [3]int = [3]int{1, 2, 3}
// 	modify(&b)
// 	fmt.Println(b)
// }

func testvariable() {
	var a *int = new(int)
	*a = 100
	fmt.Printf("*a=%d\n", *a)
	var b *[]int = new([]int)
	fmt.Printf("*b = %v\n", *b)
	(*b) = make([]int, 5, 10)
	(*b)[0] = 100
	(*b)[1] = 200
	fmt.Printf("*b = %v\n", *b)
}

func main() {
	testvariable()
}
