package main

import (
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//初始化(固定)
	gtk.Init(&os.Args)
	//用户写的代码
	//1 创建窗口
	//2 设置属性(标题，大小)
	//3 显示窗口
	win := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //带边框的顶层窗口
	win.SetTitle("go gtk")                    //设置标题
	win.SetSizeRequest(480, 320)              //设置大小
	win.Show()
	//主事件循环(固定)
	//1 让徐不结束
	//2 等待用户操作(移动窗口，鼠标点击)
	gtk.Main()
}
