package main

import (
	// "encoding/json"
	"fmt"
)

// type User struct {
// 	Username  string
// 	Sex       string
// 	Age       int
// 	AvatarUrl string
// }

// func textstruct() {
// 	var user User
// 	user.Age = 18
// 	user.Username = "user01"
// 	user.Sex = "男"
// 	user.AvatarUrl = "www.baidu.com"
// 	fmt.Printf("user.name=%s age=%d sex=%s avatar=%s\n", user.Username, user.Age, user.Sex, user.AvatarUrl)
// }

// func teststruct() {
// 	var user *User
// 	fmt.Printf("user=%v\n", user)

// 	var user01 *User = &User{}
// 	user01.Age = 18
// 	user01.Username = "user01"
// 	fmt.Printf("user01=%#v\n", user01)
// }

// type Test struct {
// 	A int32
// 	B int32
// 	C int32
// 	D int32
// }
// func teststruct() {
// 	var t Test
// 	fmt.Printf("a addr:%p\n", &t.A)
// 	fmt.Printf("b addr:%p\n", &t.B)
// 	fmt.Printf("c addr:%p\n", &t.C)
// 	fmt.Printf("d addr:%p\n", &t.D)
// }

// type Address struct {
// 	Province string
// 	City     string
// }
// type User struct {
// 	Username string
// 	Sex      string
// }
// type User01 struct {
// 	City     string
// 	username string
// 	Sex      string
// 	*Address
// }

// func teststruct() {
// 	u := &User{
// 		Username: "user03",
// 		Sex:      "man",
// 		address: Address{
// 			Province: "北京",
// 			City:     "beijing",
// 		},
// 	}
// 	fmt.Printf("user=%#v\n", u)
// }

// func teststruct(){
// 	var user User
// 	user.Username = "user01"
// 	user.Sex = "man"
// 	user.Address = &Address{	//第一种方式
// 		Province:"bj",
// 		City:"bj",
// 	}
// 	user.Province = "bj"	//第二中方式
// 	user.City = "bj"
// }

// func teststruct() {
// 	var user User01
// 	user.username = "user01"
// 	user.Sex = "man"
// 	user.City = "bj"
// 	fmt.Printf("user=%#v\n", user)
// 	user.Address = new(Address)
// 	fmt.Printf("user=%#v\n", user)
// 	user.Address.City = "bj01"
// 	fmt.Printf("user=%#v city of address:%s\n", user, user.Address.City)
// }

// type User struct {
// 	Username string `json:"username"`
// 	Sex      string `json:"sex"`
// 	Score    float32
// }

// func teststruct() {
// 	user := &User{
// 		Username: "user01",
// 		Sex:      "男",
// 		Score:    84.5,
// 	}
// 	data, date := json.Marshal(user)
// 	fmt.Printf("json str:%s\n", string(data))
// 	fmt.Println(date)
// }

// var (
// 	AllStudents []*Student
// )

// func showMenu() {
// 	fmt.Println("1.AddStudent")
// 	fmt.Println("2.ModifyStudent")
// 	fmt.Println("3.ShowAllStudent")
// 	fmt.Println("4.exit\n")
// }

// func inputStu() *Student {
// 	var (
// 		username string
// 		sex      int
// 		grade    string
// 		score    float32
// 	)
// 	fmt.Println("please input username:")
// 	fmt.Scanf("%s\n", &username)
// 	fmt.Println("please input sex:[0|1]")
// 	fmt.Scanf("%d\n", &sex)
// 	fmt.Println("please input grade:[0-6]")
// 	fmt.Scanf("%s\n", &grade)
// 	fmt.Println("please input score:[0-100]")
// 	fmt.Scanf("%f\n", &score)
// 	stu := NewStudent(username, sex, score, grade)
// 	return stu
// }

// func AddStudent() {
// 	stu := inputStu()
// 	for index, v := range AllStudents {
// 		if v.Username == stu.Username {
// 			fmt.Println("user %s success update\n\n", stu.Username)
// 			AllStudents[index] = stu
// 			return
// 		}
// 	}
// 	AllStudents = append(AllStudents, stu)
// 	fmt.Printf("user %s success insert\n\n", stu.Username)
// }

// func ModifyStudent() {
// 	stu := inputStu()
// 	for index, v := range AllStudents {
// 		if v.Username == stu.Username {
// 			AllStudents[index] = stu
// 			fmt.Printf("user %s success update\n\n", stu.Username)
// 			return
// 		}
// 		fmt.Printf("user %s is not found", v.Username)
// 	}
// }

// func ShowAllStudent() {
// 	for _, v := range AllStudents {
// 		fmt.Println("user:%s info:%#v\n", v.Username, v)
// 	}
// }

// func teststruct() {
// 	for {
// 		showMenu()
// 		var sel int
// 		fmt.Scanf("%d\n", sel)
// 		switch sel {
// 		case 1:
// 			AddStudent()
// 		case 2:
// 			ModifyStudent()
// 		case 3:
// 			ShowAllStudent()
// 		case 4:
// 			os.Exit(0)
// 		}
// 	}
// }

func main() {
	// teststruct()
	// inputStu()
	A := &Student{A: 1}
	fmt.Printf("%v", A)
}
