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

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})
	//获取控件
	entry := gtk.EntryFromObject(builder.GetObject("entry1"))
	//设置内容
	entry.SetText("123")
	//获取内容
	fmt.Println("entry.GetText()=", entry.GetText())
	//是否只读
	// entry.SetEditable(false)	//不可编辑
	//是否可见，密码模式
	// entry.SetVisibility(false)
	//"activate" 控件内部按回车键时触发
	entry.Connect("activate", func() {
		//获取内容
		fmt.Println("entry.GetText()=", entry.GetText())
	})

	//显示控件，如果是通过glade添加的控件，show即显示所有，代码布局则需要showall
	win.Show()
	//主事件循环
	gtk.Main()
}
