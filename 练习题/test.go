package main

import (
	"fmt"
)

// if justfy(a int) boll {
// 	for i:=2;i<a;i++{
// 		if a % i == 0{
// 			return false
// 		}
// 	}
// 	return true
// }
// func example1(){
// 	for i :=2;i<100;i++{
// 		if justfy(i) == true{
// 			fmt.Printf{"[%d] is prime\n",i}
// 		}
// 	}
// }

// func shuixianhua(a int) bool{
// 	first := a%10
// 	second := (a/10)%10
// 	third := (a/100)%10
// 	sum := first*first*first+second*second*second+third*third*third
// 	if sum == a{
// 		return true
// 	}
// 	return false
// }

// func example2(){
// 	for i :=100;i<1000;i++{
// 		if shuixianhua(i) == true{
// 			fmt.Printf("[%d]是水仙花数\n",i)
// 		}
// 	}
// }

func test(str string) (charCount, numCount, spaceCount, otherCount int) {
	uftChars := []rune(str)
	for i := 0; i < len(uftChars); i++ {
		if uftChars[i] >= 'a' && uftChars[i] <= 'z' || uftChars[i] >= 'A' && uftChars[i] <= 'Z' {
			charCount++
			continue
		}
		if uftChars[i] >= '0' && uftChars[i] <= '9' {
			numCount++
			continue
		}
		if uftChars[i] == ' ' {
			spaceCount++
			continue
		}
		otherCount++
	}
	return
}

func example3() {
	var str string = "sdfsd   我123+——"
	charCount, numCount, spCount, other := test(str)
	fmt.Printf("字母有%d,数字有%d,空格有%d,其他有%d", charCount, numCount, spCount, other)
}

func main() {
	// example1()
	// example2()
	example3()
}
