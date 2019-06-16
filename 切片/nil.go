package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// func testnil() {
// 	a := [5]int{1, 2, 3, 4, 5}
// 	var b []int = a[1:4]
// 	fmt.Println(b)
// }

// func testnil() {
// 	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
// 	var b []int = a[2:6]
// 	var c []int = a[3:]
// 	var d []int = a[:5]
// 	var e []int = a[:]
// 	fmt.Println(b)
// 	fmt.Println(c)
// 	fmt.Println(d)
// 	fmt.Println(e)
// }

// func testnil() {
// 	a := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
// 	fmt.Println(a)
// 	b := a[2:5]
// 	b[0] = b[0] + 10
// 	b[1] = b[1] + 10
// 	b[2] = b[2] + 10
// 	fmt.Println(a)
// 	darr := [...]int{1,2,3,4,5,6}
// 	dslice :=darr[2:5]
// 	fmt.Println("array before",darr)
// 	for i:=range dslice{
// 		dslice[i]++
// 	}
// 	fmt.Println("array after",darr)
// }

// func testnil() {
// 	numa := [3]int{78, 79, 80}
// 	nums1 := numa[:]
// 	nums2 := numa[:]
// 	fmt.Println("array before hange 1", numa)
// 	nums1[0] = 100
// 	fmt.Println("array sfter modification to slice nums1", nums1)
// 	nums2[1] = 101
// 	fmt.Println("array sfter modification to slice nums2", nums2)
// }

// func testnil() {
// 	a := make([]int, 5, 10)
// 	a[0] = 10
// 	a[1] = 20
// 	a = append(a, 30)
// 	fmt.Println(a)
// }

// func testnil() {
// 	a := [...]string{"a", "b", "c", "d", "e", "f"}
// 	b := a[1:3]
// 	fmt.Println(b, len(b), cap(b))    //[b,c]长度2，容量5
// 	b = b[:cap(b)]
// 	fmt.Println(b, len(b), cap(b))    //这样切片会让长度=容量=5
// }

// func testnil() {
// 	var a []int
// 	fmt.Println(a, len(a), cap(a))
// 	if a == nil {
// 		fmt.Println("a is nil")
// 	}
// 	a = append(a, 100)
// 	fmt.Printf("%d,%p,len:%d,cap%d\n", a, a, len(a), cap(a))
// 	a = append(a, 200)
// 	fmt.Printf("%d,%p,len:%d,cap%d\n", a, a, len(a), cap(a))
// 	a = append(a, 300)
// 	fmt.Printf("%d,%p,len:%d,cap%d\n", a, a, len(a), cap(a))
// 	a = append(a, 400)
// 	fmt.Printf("%d,%p,len:%d,cap%d\n", a, a, len(a), cap(a))
// 	a = append(a, 500)
// 	fmt.Printf("%d,%p,len:%d,cap%d\n", a, a, len(a), cap(a))
// }

// func testnil() {
// 	var a []int = []int{1, 2, 3}
// 	var b []int = []int{4, 5, 6}
// 	copy(a, b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// 		var c []int = []int{1, 2}
// 		var d []int = []int{4, 5, 6}
// 		copy(c, d)
// 		fmt.Println(c)
// 		fmt.Println(d)
// }

// func testnil() {
// 	var sa = make([]string, 5, 10)
// 	for i := 0; i < 5; i++ {
// 		sa = append(sa, fmt.Sprintf("%v", i))
// 	}
// 	fmt.Println(sa)
// }

// func testnil() {
// 	var a [5]int = [5]int{1, 4, 2, 5, 6}
// 	sort.Ints(a[:])
// 	fmt.Println(a)
// 	var b [5]string = [5]string{"ac", "ed", "sd", "fg"}
// 	sort.Strings(b[:])
// 	fmt.Println(b)
// }

var (
	length  int
	charset string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrsuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func testflag() {
	flag.IntVar(&length, "l", 16, "-l 生成密码的长度")
	flag.StringVar(&charset, "t", "num",
		`-t 制定密码生成字符集，
	num:只是用数字0~9，
	char:只使用尹文字母a-z,A-Z，
	mix:使用数字和字母，
	advance:使用数字、字母以及特殊字符`)
	flag.Parse()
}

func generatePasswd() string {
	var passwd []byte = make([]byte, length, length)
	var sourceStr string
	if charset == "num" {
		sourceStr = NumStr
	} else if charset == "char" {
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr)
	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr)
	} else {
		sourceStr = NumStr
	}
	fmt.Println("source", sourceStr)
	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return string(passwd)
}

func main() {
	// testnil()
	rand.Seed(time.Now().UnixNano())
	testflag()
	// fmt.Printf("length:%d charset:%s\n", length, charset)
	passwd := generatePasswd()
	fmt.Println(passwd)
}
