package main

import (
	"fmt"
)

// type Animal interface {
// 	Talk()
// 	Eat()
// 	Name() string
// }
// type Dog struct {
// }
// type Pig struct {
// }

// func (d Dog) Talk() {
// 	fmt.Println("汪汪")
// }
// func (d Dog) Eat() {
// 	fmt.Println("吃骨头")
// }
// func (d Dog) Name() string {
// 	fmt.Println("旺财")
// 	return "旺财"
// }
// func (p Pig) Talk() {
// 	fmt.Println("哼哼")
// }
// func (d Pig) Eat() {
// 	fmt.Println("吃肉")
// }
// func (d Pig) Name() string {
// 	fmt.Println("佩奇")
// 	return "佩奇"
// }

// func test() {
// 	var b Dog
// 	var a Animal
// 	a = b
// 	a.Name()
// 	a.Eat()
// 	a.Talk()
// 	var p Pig
// 	a = p
// 	a.Name()
// 	a.Eat()
// 	a.Talk()
// }

// type Employer interface {
// 	CalcSalary() float32
// }
// type Programer struct {
// 	name  string
// 	base  float32
// 	extra float32
// }

// func (p Programer) CalcSalary() float32 {
// 	return p.base
// }
// func NewProgramer(name string, base float32, extra float32) Programer {
// 	return Programer{
// 		name:  name,
// 		base:  base,
// 		extra: extra,
// 	}
// }

// type Sale struct {
// 	name  string
// 	base  float32
// 	extra float32
// }

// func (p Sale) CalcSalary() float32 {
// 	return p.base + p.extra*p.base*0.5
// }
// func NewSale(name string, base, extra float32) Sale {
// 	return Sale{
// 		name:  name,
// 		base:  base,
// 		extra: extra,
// 	}
// }
// func calcAll(e []Employer) float32 {
// 	var cost float32
// 	for _, v := range e {
// 		cost = cost + v.CalcSalary()
// 	}
// 	return cost
// }
// func test() {
// 	p1 := NewProgramer("搬砖1", 1500.0, 0)
// 	p2 := NewProgramer("搬砖2", 1500.0, 0)
// 	p3 := NewProgramer("搬砖3", 1500.0, 0)
// 	s1 := NewSale("销售1", 800.0, 2.5)
// 	s2 := NewSale("销售2", 800.0, 2.5)
// 	s3 := NewSale("销售3", 800.0, 2.5)
// 	var employList []Employer
// 	employList = append(employList, p1)
// 	employList = append(employList, p2)
// 	employList = append(employList, p3)
// 	employList = append(employList, s1)
// 	employList = append(employList, s2)
// 	employList = append(employList, s3)
// 	cost := calcAll(employList)
// 	fmt.Printf("这个月人力成本:%f\n", cost)
// }

// func test() {
// 	var a interface{}
// 	var b int = 100
// 	a = b
// 	fmt.Printf("%T %v\n", a, a)
// 	var c string = "hi"
// 	a = c
// 	fmt.Printf("%T %v\n", a, a)
// 	var d map[string]int = make(map[string]int, 10)
// 	d["abc"] = 100
// 	d["aaa"] = 30
// 	a = d
// 	fmt.Printf("%T %v\n", a, a)
// }

// func describe(a interface{}) {
// 	fmt.Printf("type=%T %v\n", a, a)
// }

// type Student struct {
// 	Name string
// 	Sex  int
// }

// func test() {
// 	a := 99
// 	describe(a)
// 	b := "hello"
// 	describe(b)
// 	var stu Student = Student{
// 		Name: "user01",
// 		Sex:  1,
// 	}
// 	describe(stu)
// }

// func describe(a interface{}) {
// 	s, ok := a.(int)
// 	if ok {
// 		fmt.Println(s)
// 		return
// 	}
// 	str, ok := a.(string)
// 	if ok {
// 		fmt.Println(str)
// 		return
// 	}
// 	f, ok := a.(float32)
// 	if ok {
// 		fmt.Println(f)
// 		return
// 	}
// 	fmt.Println("输入错误")
// 	return

// }
// func test() {
// 	var a int = 100
// 	describe(a)
// 	var b string = "hi"
// 	describe(b)
// }
// func testSwitch(a interface{}) {
// 	switch v := a.(type) {
// 	case string:
// 		fmt.Printf("a is string,value:%v\n", v)
// 	case int:
// 		fmt.Printf("a is int,value:%v\n", v)
// 	case float32:
// 		fmt.Printf("a is float32,value:%v\n", v)
// 	default:
// 		fmt.Println("not support type\n")
// 	}
// }
// func test() {
// 	var a int = 100
// 	testSwitch(a)
// 	var b string = "hello"
// 	testSwitch(b)
// }

// type Animal interface {
// 	Talk()
// }
// type Bu interface {
// 	Bu()
// }
// type Dog struct {
// }

// func (d Dog) Talk() {
// 	fmt.Println("汪汪汪")
// }
// func (d Dog) Bu() {
// 	fmt.Println("狗是哺乳动物")
// }
// func test() {
// 	var d Dog
// 	var a Animal
// 	a = d
// 	a.Talk()
// 	var b Bu
// 	b = d
// 	b.Bu()
// }

type Animal interface {
	Eat()
}
type Describle interface {
	Describled()
}
type AdvanceAnimal interface {
	Animal
	Describle
}
type Dog struct {
}

func (d Dog) Eat() {
	fmt.Println("狗吃屎")
}
func (d Dog) Describled() {
	fmt.Println("狗")
}
func test() {
	var d Dog
	var a AdvanceAnimal
	a = d
	a.Eat()
	a.Describled()
}
func main() {
	test()
}
