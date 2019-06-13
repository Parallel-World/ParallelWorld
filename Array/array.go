package main

import "fmt"

// func testArray() {
// 	var a [5]int
// 	a[0] = 200
// 	a[1] = 300
// 	fmt.Println(a)
// }

// func testArray(){
// 	a :=[5]int{3:100,4:300}
// 	for i :=0;i<len(a); i++{
// 		fmt.Printf("a[%d]=%d\n",i,a[i])
// 	}
// }

// func testArray() {
// 	a := [3]int{1, 2, 3}
// 	for index, value := range a {
// 		fmt.Printf("a[%d]=%d\n", index, value)
// 	}
// }

// func testArray() {
// 	var a [3][2]int
// 	a[0][0] = 10
// 	a[0][1] = 20
// 	a[1][0] = 30
// 	a[1][1] = 40
// 	a[2][0] = 50
// 	a[2][1] = 60
// 	fmt.Println(a)

// for i := 0; i < 3; i++ {
// 	for j := 0; j < 2; j++ {
// 		fmt.Printf("%d ", a[i][j])
// 	}
// }
// for i, val1 := range a {
// 	for j, val2 := range val1 {
// 		fmt.Printf("(%d,%d)=%d ", i, j, val2)
// 	}
// }
// }

// func testArray() {
// 	var a [3]int
// 	a[0] = 0
// 	a[1] = 1
// 	a[2] = 2
// 	b := a
// 	b[0] = 1000
// 	fmt.Println(a, b)
// }

func testArray() {
	a := [3]int{10, 20, 30}
	modify(a)
	fmt.Print(a)
}
func modify(b [3]int) {
	b[0] = 1000
	return
}

func main() {
	testArray()
}
