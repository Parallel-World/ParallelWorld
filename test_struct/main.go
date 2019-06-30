package main

import (
	"fmt"
	"os"
)

var (
	studenMar = &StudentMgr{}
)

func showMenu() {
	fmt.Println("1.add student")
	fmt.Println("2.modify student")
	fmt.Println("3.show all student")
	fmt.Println("4.exit\n")
}
func inputStudent() *Student {
	var (
		username string
		sex      int
		grade    string
		score    float32
	)
	fmt.Println("Please input username:")
	fmt.Scanf("%s\n", &username)
	fmt.Println("Please input sex:[0|1]")
	fmt.Scanf("%d\n", &sex)
	fmt.Println("Please input grade:[0-6]")
	fmt.Scanf("%s\n", &grade)
	fmt.Println("Please input score:[0-100]")
	fmt.Scanf("%f\n", &score)
	stu := NewStudent(username, sex, score, grade)
	return stu
}
func main() {
	for {
		showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			stu := inputStudent()
			studenMar.AddStudent(stu)
		case 2:
			stu := inputStudent()
			studenMar.ModifyStudent(stu)
		case 3:
			studenMar.ShowAllstudent()
		case 4:
			os.Exit(0)
		}
	}
}
