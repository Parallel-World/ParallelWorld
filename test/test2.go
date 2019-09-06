package main

import "fmt"

// package main

// import "fmt"

// // func testArray() {
// // 	a := [...]int{1, 2, 3, 4, 5, 8, 4, 4, 2}
// // 	for i := 0; i < len(a); i++ {
// // 		for j := i + 1; j < len(a); j++ {
// // 			if a[i]+a[j] == 8 {
// // 				fmt.Printf("(%d,%d)的和等于8\n", i, j)
// // 			}
// // 		}
// // 	}
// // }

// func test(a *int) {
// 	fmt.Println(a, *a)
// 	*a = 20
// 	return
// }

// func main() {
// 	// testArray()
// 	a := 10
// 	fmt.Println(&a, a)
// 	test(&a)
// 	fmt.Println(&a, a)
// }

// func test() {
// 	darr := [...]int{1, 2, 3, 4, 5, 6}
// 	dslice := darr[2:5]
// 	fmt.Println(dslice)
// 	for i, v := range dslice {
// 		dslice[i]++
// 		fmt.Println(i, v)
// 	}
// 	fmt.Println(dslice)
// }

// func test() {
// 	a := [...]string{"a", "b", "c", "d", "e", "f"}
// 	b := a[1:3]
// 	fmt.Println(b, len(b), cap(b))
// 	b = b[:cap(b)]
// 	fmt.Println(b, len(b), cap(b))
// }

// type User struct {
// 	Username string
// 	Sex      string
// 	Age      int
// 	int
// 	string
// }

// func test() {
// 	var u User
// 	u.Username = "user01"
// 	u.Sex = "man"
// 	u.int = 100
// 	u.string = "hello"
// 	fmt.Printf("user = %#v\n", u)
// }

type Student struct{
	Name string
	Next *Student
	Prev *Student
}

type


func main() {
	test()
}
