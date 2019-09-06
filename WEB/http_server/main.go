package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
)

// func sayHello(w http.ResponseWriter, r *http.Request) { //w表示写，r表示读取的信息
// 	r.ParseForm() //解析参数，默认不解析
// 	fmt.Fprintf(w, "%v\n", r.Form)
// 	fmt.Fprintf(w, "path:%s\n", r.URL.Path)
// 	fmt.Fprintf(w, "schema:%s\n", r.URL.Scheme)
// 	fmt.Fprintf(w, "hello world\n")
// }
// func test() {
// 	http.HandleFunc("/", sayHello)           //设置访问的路由，第二个参数为函数名
// 	err := http.ListenAndServe(":8080", nil) //设置监听的端口，阻塞
// 	if err != nil {
// 		fmt.Printf("http server failed,err:%v\n", err)
// 	}
// }

// func login(w http.ResponseWriter, r *http.Request) {
// 	method := r.Method
// 	if method == "GET" {
// 		t, err := template.ParseFiles("./login.html")
// 		if err != nil {
// 			fmt.Fprintf(w, "load login.html failed")
// 			return
// 		}
// 		t.Execute(w, nil)
// 	} else if method == "POST" {
// 		r.ParseForm()
// 		username := r.FormValue("username")
// 		pwd := r.FormValue("password")
// 		fmt.Printf("username:%s\n", username)
// 		fmt.Printf("password:%s\n", pwd)
// 		if username == "admin" && pwd == "admin" {
// 			fmt.Fprintf(w, "username:%s login success\n", r.FormValue("username"))
// 		} else {
// 			fmt.Fprintf(w, "username:%s login failed\n", r.FormValue("username"))
// 		}
// 	}
// }

type Address struct {
	City     string
	Province string
}
type UserInfo struct {
	Name    string
	Sex     string
	Age     int
	Address Address
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./login.html")
	if err != nil {
		fmt.Fprintf(w, "load login.html failed,err:%v", err)
		return
	}
	var userlist []*UserInfo
	for i := 0; i < 10; i++ {
		// 结构体
		user := UserInfo{
			Name: fmt.Sprintf("Mary%d", rand.Intn(80)),
			Sex:  "男",
			Age:  rand.Intn(50),
			Address: Address{
				City:     "北京",
				Province: "北京市",
			},
		}
		userlist = append(userlist, &user)
	}
	err = t.Execute(w, userlist)
	//map
	// m := make(map[string]interface{})
	// m["Name"] = "Mary"
	// m["Sex"] = "男"
	// m["Age"] = 18
	// err = t.Execute(w, m)
	if err != nil {
		fmt.Printf("execute template failed,err:%v\n", err)
	}
	// t.Execute(os.Stdout, m) //显示在终端

}

func test() {
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("listen server failed,err:%v\n", err)
		return
	}
}
func main() {
	test()
}
