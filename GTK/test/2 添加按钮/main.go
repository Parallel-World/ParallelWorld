package main

import (
	"os"

	"github.com/mattn/go-gtk/gtk"
)

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
	//主事件循环
	gtk.Main()
}
