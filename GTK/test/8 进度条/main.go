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
	//进度条控件
	pg := gtk.ProgressBarFromObject(builder.GetObject("progressbar1"))
	//设置进度0.0~1.0
	pg.SetFraction(0.5)
	//获取进度并打印
	fmt.Println("value=", pg.GetFraction())
	//设置文本内容
	pg.SetText("50%")

	//显示控件，如果是通过glade添加的控件，show即显示所有，代码布局则需要showall
	win.Show()
	//主事件循环
	gtk.Main()
}
