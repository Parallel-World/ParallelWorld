package main

import "fmt"

// func testmap() {
// 	var a map[string]int
// 	fmt.Printf("a:%v\n", a)
// 	// a["stu01"] = 100
// 	if a == nil {
// 		a = make(map[string]int, 16)
// 		fmt.Printf("a=%v\n", a)
// 		a["stu01"] = 1000
// 		fmt.Printf("a=%v\n", a)
// 	}
// }

// func testmap() {
// 	a := make(map[string]int)
// 	a["steve"] = 12000
// 	a["jamie"] = 15000
// 	a["mike"] = 9000
// 	fmt.Println("a map contents:", a)
// }

// func testmap() {
// 	a := map[string]int{
// 		"steve": 12000,
// 		"jamie": 15000,
// 	}
// 	a["mike"] = 9000
// 	fmt.Println("a map contents:", a)
// }

// func testmap() {
// 	a := map[string]int{
// 		"steve": 12000,
// 		"jamie": 15000,
// 	}
// 	a["mike"] = 9000
// 	b := "jamie"
// 	fmt.Println("Salary of", b, "is", a[b])
// }

// func testmap() {
// 	a := map[string]int{
// 		"stu01": 1000,
// 		"stu02": 2000,
// 	}
// 	b := "joe"
// 	value, ok := a[b]
// 	if ok == false {
// 		fmt.Printf("key %s is not exist\n", b)
// 	} else {
// 		fmt.Printf("key %s is %d\n", b, value)
// 	}
// }

// func testmap() {
// 	a := map[string]int{
// 		"stece": 1000,
// 		"jamie": 2000,
// 	}
// 	fmt.Println("all items of a map")
// 	for key, value := range a {
// 		fmt.Printf("[%s]=%d\n", key, value)
// 	}
// }

// func testmap() {
// 	a := map[string]int{
// 		"steve": 1000,
// 		"jamie": 2000,
// 	}
// 	fmt.Println("map before deletion", a)
// 	delete(a, "steve")
// 	fmt.Println("map after deletion", a)
// }

// func testmap() {
// 	a := map[string]int{
// 		"steve": 1000,
// 		"jamie": 2000,
// 	}
// 	fmt.Println("length is", len(a))
// }

// func testmap() {
// 	a := map[string]int{
// 		"steve": 1000,
// 		"jamie": 2000,
// 	}
// 	fmt.Println("origin map", a)
// 	b := a
// 	b["steve"] = 1800
// 	fmt.Println("a map changed", a)
// }

// func testmap() {
// 	var a map[string]int = make(map[string]int, 10)
// 	for i := 0; i < 10; i++ {
// 		key := fmt.Sprintf("key%d", i)
// 		a[key] = i
// 	}
// 	var keys []string = make([]string, 0, 10)
// 	for key, value := range a {
// 		fmt.Printf("key:%s = %d\n", key, value)
// 		keys = append(keys, key)
// 	}
// 	sort.Strings(keys)
// 	for _, value := range keys {
// 		fmt.Printf("key:%s val:%d\n", value, a[value])
// 	}
// }

// func testmap() {
// 	var mapSlice []map[string]int
// 	mapSlice = make([]map[string]int, 3, 10)
// 	fmt.Println("before map init")
// 	for index, value := range mapSlice {
// 		fmt.Printf("index:%d value:%v\n", index, value)
// 	}
// 	fmt.Println()
// 	mapSlice[0] = make(map[string]int, 10)
// 	mapSlice[0]["a"] = 1000
// 	mapSlice[0]["b"] = 2000
// 	fmt.Println("after map init")
// 	for index, value := range mapSlice {
// 		fmt.Printf("index:%d value:%v\n", index, value)
// 	}
// }

// func testmap() {
// 	var s map[string][]int
// 	s = make(map[string][]int, 16)
// 	key := "stu01"
// 	value, ok := s[key]
// 	if !ok {
// 		s[key] = make([]int, 0, 16)
// 	}
// 	value = append(value, 100)
// 	value = append(value, 200)
// 	s[key] = value
// 	fmt.Printf("map:%v\n", s)
// }

// func test(str string) map[string]int {

// 	var result map[string]int = make(map[string]int, 128)
// 	words := strings.Split(str, " ")
// 	for _, v := range words {
// 		count, ok := result[v]
// 		if !ok {
// 			result[v] = 1
// 		} else {
// 			result[v] = count + 1
// 		}
// 	}
// 	return result
// }

// func test() {
// 	var a interface{}
// 	var b int = 100
// 	var c float32 = 1.2
// 	var d string = "hello"
// 	a = b
// 	fmt.Println("a=", a)
// 	a = c
// 	fmt.Println("a=", a)
// 	a = d
// 	fmt.Println("a=", a)
// }

// func test() {
// 	var stuMap map[int]map[string]interface{}
// 	stuMap = make(map[int]map[string]interface{}, 16)
// 	var id = 1
// 	var name = "stu01"
// 	var score = 79.3
// 	var age = 18
// 	value, ok := stuMap[id]
// 	if !ok {
// 		value = make(map[string]interface{}, 8)
// 	}
// 	value["id"] = id
// 	value["name"] = name
// 	value["score"] = score
// 	value["age"] = age
// 	stuMap[id] = value
// 	fmt.Printf("stuMap:%v\n", stuMap)
// 	fmt.Printf("stuMap:%#v\n", stuMap)
// 	for i := 0; i < 10; i++ {
// 		value, ok := stuMap[i]
// 		if !ok {
// 			value = make(map[string]interface{}, 8)
// 		}
// 		value["name"] = fmt.Sprintf("stu%d", i)
// 		value["id"] = i
// 		value["score"] = rand.Float32() * 100.0
// 		value["age"] = rand.Intn(100)
// 		stuMap[i] = value
// 	}
// 	fmt.Println(stuMap)
// 	for k, v := range stuMap {
// 		fmt.Printf("id=%d stu info=%v\n", k, v)
// 		fmt.Printf("id=%d stu info=%#v\n", k, v)
// 	}
// }

// var (
// 	coins = 50
// 	users = []string{
// 		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
// 		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
// 	}
// 	distribution = make(map[string]int, len(users))
// )

// func test() {
// 	for i := 0; i < len(users); i++ {
// 		v := 0
// 		for j := 0; j < len(users[i]); j++ {
// 			q := string(users[i][j])
// 			if q == "a" {
// 				v = v + 1
// 			} else if q == "A" {
// 				v = v + 1
// 			} else if q == "e" {
// 				v = v + 1
// 			} else if q == "E" {
// 				v = v + 1
// 			} else if q == "i" {
// 				v = v + 2
// 			} else if q == "I" {
// 				v = v + 1
// 			} else if q == "o" {
// 				v = v + 3
// 			} else if q == "O" {
// 				v = v + 3
// 			} else if q == "u" {
// 				v = v + 5
// 			} else if q == "U" {
// 				v = v + 5
// 			}
// 		}
// 		distribution[users[i]] = v
// 		coins = coins - v
// 	}
// 	fmt.Printf("每个人分到%v个金币，剩余%d个金币", distribution, coins)
// }

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func calcCoin(username string) int {
	var sum int = 0
	for _, char := range username {
		switch char {
		case 'a', 'A', 'e', 'E':
			sum = sum + 1
		case 'i', 'I':
			sum = sum + 2
		case 'o', 'O':
			sum = sum + 3
		case 'u', 'U':
			sum = sum + 5
		}
	}
	return sum
}

func distribution_coin() int {
	var left int = coins
	for _, username := range users {
		allCoin := calcCoin(username)
		left = left - allCoin
		value, ok := distribution[username]
		if !ok {
			distribution[username] = allCoin
		} else {
			distribution[username] = value + allCoin
		}
	}
	return left
}

func main() {
	// testmap()
	// var str = "how do you do?"
	// result := test(str)
	// fmt.Printf("result:%v\n", result)
	// test()
	left := distribution_coin()
	for username, coin := range distribution {
		fmt.Printf("user:%s have %d coins\n", username, coin)
	}
	fmt.Println("left coin:", left)
}
