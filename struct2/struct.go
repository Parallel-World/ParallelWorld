package main

import "fmt"

type People struct {
	Name    string
	Country string
}

func (p People) Print() {
	fmt.Printf("name=%s country=%s\n", p.Name, p.Country)
}
func (p People) Set(name, country string) {
	p.Name = name
	p.Country = country
}

func test() {
	var p1 People = People{
		Name:    "pe01",
		Country: "china",
	}
	p1.Print()
	p1.Set("pe02", "Ch")
	p1.Print()
}

// type Integer int

// func (i Integer) Print() {
// 	fmt.Println(i)
// }

// func test() {
// 	var a Integer
// 	a = 1000
// 	a.Print()

// 	var b int = 200
// 	a = Integer(b)
// 	a.Print()
// }

// type People struct {
// 	Name    string
// 	Country string
// }

// func (p People) Print() {
// 	fmt.Printf("name=%s country=%s\n", p.Name, p.Country)
// }
// func (p *People) Set2(name, country string) {
// 	p.Name = name
// 	p.Country = country
// }
// func test() {
// 	var p1 People = People{
// 		Name:    "pe01",
// 		Country: "china",
// 	}
// 	p1.Print()
// 	(&p1).Set2("pe02", "Ch")
// 	p1.Print()
// }

// type Animal struct {
// 	Name string
// 	Sex  string
// }

// func (a *Animal) Talk() {
// 	fmt.Printf("i,talk,i'm %s\n", a.Name)
// }

// type Dog struct {
// 	Feet string
// 	*Animal
// }

// func (d *Dog) Eat() {
// 	fmt.Println("dog is eat")
// }
// func test() {
// 	var d *Dog = &Dog{
// 		Feet: "four feet",
// 		Animal: &Animal{
// 			Name: "dog",
// 			Sex:  "xiong",
// 		},
// 	}
// 	d.Eat()
// 	d.Talk()
// }

// type Dog struct {
// 	Feet string
// 	*Animal1
// 	*Animal2
// }

// type Animal1 struct {
// 	Name string
// 	Sex  string
// }

// func (a *Animal1) Talk() {
// 	fmt.Printf("i,talk,i'm %s\n", a.Name)
// }

// type Animal2 struct {
// 	Name string
// 	Sex  string
// }

// func (p *Animal2) Talk() {
// 	fmt.Println("Animal2")
// }
// func test() {
// 	var d *Dog = &Dog{
// 		Feet: "four feet",
// 		Animal1: &Animal1{
// 			Name: "dog",
// 			Sex:  "xiong",
// 		},
// 	}
// 	d.Animal1.Talk()
// 	d.Animal2.Talk()
// }
func main() {
	test()
}
