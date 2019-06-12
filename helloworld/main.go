package main

import "fmt"

func main() {
	// s1 := "hello"
	// byteArray := []byte(s1)  // [h e l l o]
	// fmt.Println(byteArray)
	// s2 := ""
	// for i:=len(byteArray)-1;i>=0; i--{
	//i 是 4 3 2 1 0
	// byteArray[i] o l l e h (字符)
	// 	s2 = s2 + string(byteArray[i])  // 102--> "h"
	// 	fmt.Println(s2)
	// }
	// 	fmt.Println(s2)
	// var a = [...]int{1, 3, 5, 7, 8}
	// // var b int
	// // b = 0
	// fmt.Println(a)
	// sum := 0
	// for _, v := range a {
	// 	sum = sum + v
	// }
	// fmt.Println(sum)
	// a := "平行世界"
	// a_b := []rune(a)
	// for i := 0; i < len(a_b)/2; i++ {
	// 	a_b[i], a_b[len(a_b)-i-1] = a_b[len(a_b)-i-1], a_b[i]
	// }
	// a = string(a_b)
	// println(a)
	//
	// start := time.Now().UnixNano()
	// for i := 0; i < 10; i++ {
	// 	time.Sleep(time.Millisecond)
	// }
	// end := time.Now().UnixNano()
	// cost := (end - start) / 1000000
	// fmt.Printf("%v\n", cost)
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
