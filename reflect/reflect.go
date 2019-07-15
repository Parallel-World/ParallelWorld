package main

import (
	"fmt"
	"reflect"
)

// func reflect_example(a interface{}) {
//     t := reflect.TypeOf(a)
//     fmt.Printf("type of a is:%v\n", t)
//     k := t.Kind()
//     switch k {
//     case reflect.Int64:
//         fmt.Printf("a is int64")
//     case reflect.String:
//         fmt.Printf("a is string")
//     }
// }
// func test() {
// 	var x float32 = 3.4
// 	reflect_example(x)
// }

// func reflect_valur(a interface{}) {
// 	v := reflect.ValueOf(a)
// 	k := v.Kind()
// 	switch k {
// 	case reflect.Int64:
// 		fmt.Printf("a is int64, value is:%d", v.Int())
// 	case reflect.Float64:
// 		fmt.Printf("a is Float64, value is:%f", v.Float())
// 	}
// }
// func test() {
// 	var x float64 = 3.4
// 	reflect_valur(x)
// }

// func reflect_set_valur(a interface{}) {
// 	v := reflect.ValueOf(a)
// 	k := v.Kind()
// 	switch k {
// 	case reflect.Int64:
// 		v.SetInt(100)
// 		fmt.Printf("a is int64, value is:%d", v.Int())
// 	case reflect.Float64:
// 		v.SetFloat(6.8)
// 		fmt.Printf("a is Float64, value is:%f", v.Float())
// 	case reflect.Ptr:
// 		fmt.Printf("set a to 6.8\n")
// 		v.Elem().SetFloat(6.8)
// 	}
// }
// func test() {
// 	var x float64 = 3.4
// 	reflect_set_valur(&x)
// 	fmt.Printf("x value is %v\n", x)
// }
// type Student struct {
// 	Name string
// 	Sex  int
// 	Age  int
// }

// func test() {
// 	var s Student
// 	v := reflect.ValueOf(s)
// 	t := v.Type()
// 	kind := t.Kind()
// 	switch kind {
// 	case reflect.Int64:
// 		fmt.Printf("s is int64\n")
// 	case reflect.Float32:
// 		fmt.Printf("s is Float32\n")
// 	case reflect.Struct:
// 		fmt.Printf("s is struct\n")
// 		fmt.Printf("field num of s is %d\n", v.NumField())
// 		for i := 0; i < v.NumField(); i++ {
// 			field := v.Field(i)
// 			fmt.Printf("name:%s type:%v value:%v\n", t.Field(i).Name, field.Type(), field.Interface())
// 		}
// 	default:
// 		fmt.Println("default")
// 	}
// }

// type Student struct {
// 	Name string
// 	Sex  int
// 	Age  int
// }

// func test() {
// 	var s Student
// 	v := reflect.ValueOf(&s)
// 	v.Elem().Field(0).SetString("stu01")
// 	v.Elem().FieldByName("Sex").SetInt(2)
// 	v.Elem().FieldByName("Age").SetInt(15)
// 	fmt.Printf("s: %#v\n", s)
// }
// type Student struct {
// 	Name  string
// 	Sex   int
// 	Age   int
// 	Score float32
// }

// func (s *Student) SetName(name string) {
// 	s.Name = name
// }
// func (s *Student) Print(){
// 	fmt.Printf("%#v\n",s)
// }
// func test() {
// 	var s Student
// 	s.SetName("xxx")
// 	v := reflect.ValueOf(&s)
// 	t := v.Type()
// 	fmt.Printf("struct student have %d methods\n", t.NumMethod())
// 	for i := 0; i < t.NumMethod(); i++ {
// 		method := t.Method(i)
// 		fmt.Printf("struct %d method , name:%s type :%v\n", i, method.Name, method.Type)
// 	}
// }
// type Student struct {
// 	Name  string
// 	Sex   int
// 	Age   int
// 	Score float32
// }

// func (s *Student) SetName(name string) {
// 	s.Name = name
// }
// func (s *Student) Print() {
// 	fmt.Printf("%#v\n", s)
// }
// func test() {
// 	var s Student
// 	s.SetName("xxx")
// 	v := reflect.ValueOf(&s)
// 	m1 := v.MethodByName("Print")
// 	var args []reflect.Value
// 	m1.Call(args)
// 	m2 := v.MethodByName("SetName")
// 	var args2 []reflect.Value
// 	name := "stu01"
// 	nameVal := reflect.ValueOf(name)
// 	args2 = append(args2, nameVal)
// 	m2.Call(args2)
// 	m1.Call(args)
// }

type Student struct {
	Name  string `json:"name" db:"name2"`
	Sex   int
	Age   int
	Score float32
}

func (s *Student) SetName(name string) {
	s.Name = name
}
func (s *Student) Print() {
	fmt.Printf("%#v\n", s)
}
func test() {
	var s Student
	s.SetName("xxx")
	v := reflect.ValueOf(&s)
	t := v.Type()
	field0 := t.Elem().Field(0)
	fmt.Printf("tag json=%s\n", field0.Tag.Get("json"))
	fmt.Printf("tag db=%s\n", field0.Tag.Get("db"))
}

func main() {
	test()
}
