package main

import "fmt"

// func testnil() {
// 	a := [5]int{1, 2, 3, 4, 5}
// 	var b []int = a[1:4]
// 	fmt.Println(b)
// }

func testnil() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var b []int = a[2:6]
	var c []int = a[3:]
	var d []int = a[:5]
	var e []int = a[:]
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func main() {
	testnil()
}
