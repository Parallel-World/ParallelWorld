package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//加载glade文件
	builder := gtk.NewBuilder()
	builder.AddFromFile("test.glade")
	//获取glade上的控件
	win := gtk.WindowFromObject(builder.GetObject("window1")) //控件名称
	b1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	b2 := gtk.ButtonFromObject(builder.GetObject("button2"))
	//信号处理
	b1.Clicked(func() {
		fmt.Println("按钮1")
	})
	b2.Clicked(func() {
		fmt.Println("按钮2")
	})
	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})
	//显示控件，如果是通过glade添加的控件，show即显示所有，代码布局则需要showall
	win.Show()
	//主事件循环
	gtk.Main()
}
