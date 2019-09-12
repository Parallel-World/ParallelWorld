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
	win := gtk.WindowFromObject(builder.GetObject("windows1")) //控件名称
	win.SetTitle("Parallel")                                   //设置标题
	win.SetIconFromFile("../Image/1.png")                      //设置图标
	//设置窗口最小大小
	win.SetSizeRequest(480, 320)
	//获取窗口大小
	var w, h int
	win.GetSize(&w, &h) //获取大小
	fmt.Printf("w = %d,h = %d\n", w, h)
	//让窗口居中弹出
	win.SetPosition(gtk.WIN_POS_CENTER)
	//关闭窗口
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	// //显示控件
	win.Show()
	//主事件循环
	gtk.Main()
}
