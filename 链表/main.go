package main

import "fmt"

// type Student struct {
// 	Name  string
// 	Age   int
// 	Score float64
// 	next  *Student
// }

// func trans(p *Student) {
// 	for p != nil {
// 		fmt.Println(*p)
// 		p = p.next
// 	}
// 	fmt.Println()
// }

// func testz() {
// 	var head Student
// 	head.Name = "zhang"
// 	head.Age = 18
// 	head.Score = 100
// 	var tail = &head
// 	for i := 0; i < 5; i++ {
// 		stu := &Student{
// 			Name:  fmt.Sprintf("stu%d", i),
// 			Age:   rand.Intn(100),
// 			Score: rand.Float64() * 100,
// 		}
// 		tail.next = stu
// 		tail = stu
// 	}
// 	trans(&head)
// }

// func testn() {
// 	var head *Student = &Student{}
// 	head.Name = "hua"
// 	head.Age = 18
// 	head.Score = 100
// 	for i := 0; i < 5; i++ {
// 		stu := Student{
// 			Name:  fmt.Sprintf("stu%d", i),
// 			Age:   rand.Intn(100),
// 			Score: rand.Float64() * 100,
// 		}
// 		stu.next = head
// 		head = &stu
// 	}
// 	trans(head)
// }

// func main() {
// 	testz()
// 	testn()
// }

type Student struct {
	Name  string
	Age   int
	Score float64
	left  *Student
	right *Student
}

func trans(root *Student) {
	if root == nil {
		return
	}
	fmt.Println(root)
	trans(root.left)
	trans(root.right)
}

func main() {
	var root *Student = new(Student)
	root.Name = "stu01"
	root.Age = 19
	root.Score = 100

	var left1 *Student = new(Student)
	left1.Name = "stu02"
	left1.Age = 19
	left1.Score = 100

	root.left = left1

	var right1 *Student = new(Student)
	right1.Name = "stu04"
	right1.Age = 19
	right1.Score = 100

	root.right = right1

	var left2 *Student = new(Student)
	left2.Name = "stu03"
	left2.Age = 19
	left2.Score = 100

	left1.left = left2

	trans(root)
}
