package main

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

type UserMainController struct {
	beego.Controller
}

func (this *UserMainController) Get() {
	this.Ctx.WriteString("hello")
}
func main() {
	beego.Router("/", &MainController{})
	beego.Router("/user", &UserMainController{})
	beego.Run()
}
