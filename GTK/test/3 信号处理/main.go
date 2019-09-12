package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func HandleSignal(ctx *glib.CallbackContext) {
	fmt.Println("================")
	arg := ctx.Data()        //获取用户传递的参数，是空接口类型
	data, ok := arg.(string) //类型断言，判断是不是string
	if ok {
		fmt.Printf("按钮1被按下：%s\n", data)
	}
}
func main() {
	//初始化
	gtk.Init(&os.Args)
	//1创建主窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	//2设置窗口属性(大小、标题)
	win.SetTitle("Parallel")
	win.SetSizeRequest(480, 320)
	//3创建容器控件(固定布局，任意布局)
	layout := gtk.NewFixed()
	//4布局添加到窗口上
	win.Add(layout)
	//5创建按钮
	b1 := gtk.NewButtonWithLabel("^_^")
	b2 := gtk.NewButtonWithLabel("@_@")
	b2.SetSizeRequest(100, 100) //设置按钮大小
	//6按钮添加到布局
	layout.Put(b1, 0, 0) //b1布局的第一个参数是坐标X，第二个是Y，从左上角开始
	layout.Put(b2, 100, 100)
	//7显示控件
	//如果有多个控件，使用show，需要每一个控件都要show
	// win.Show()
	// layout.Show()
	// b1.Show()
	win.ShowAll() //所有控件都显示
	//8信号处理
	//按钮按下触发的信号："clicked"
	str := "are u ojbk?"
	//告诉系统只要按下按钮，自动调用HandleSignal,Str是给HandleSignal传递的参数
	//Connect()只会调用过一次，告诉系统一个规则
	b1.Connect("clicked", HandleSignal, str)
	/*
		//处理函数可以是匿名函数，推荐写法
		b2.Connect("clicked", func() {
			fmt.Printf("按钮2被按下：%s\n", str)
		})
	*/
	b2.Clicked(func() {
		fmt.Printf("按钮2被按下：%s\n", str)
	}) //这种写法和上面等价

	//按窗口的关闭按钮，触发"destroy"
	win.Connect("destroy", func() {
		fmt.Println("++++++++++++")
		gtk.MainQuit() //gtk程序关闭
	})

	//主事件循环
	gtk.Main()
	fmt.Println("窗口已关闭")
}
