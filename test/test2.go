package main

import "fmt"

func testArray() {
	a := [...]int{1, 2, 3, 4, 5, 8, 4, 4, 2}
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 8 {
				fmt.Printf("(%d,%d)的和等于8\n", i, j)
			}
		}
	}
}

func main() {
	testArray()
}
