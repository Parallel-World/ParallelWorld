package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gdk"
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
	//获取标签控件
	labelOne := gtk.LabelFromObject(builder.GetObject("label1"))
	labelTwo := gtk.LabelFromObject(builder.GetObject("label2"))
	labelOne.SetText("Parallel")                               //设置标签内容
	labelOne.ModifyFontSize(20)                                //设置标签大小
	labelOne.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white")) //设置字体颜色
	labelTwo.SetText("are u ok?")
	//获取内容
	str := labelTwo.GetText()
	fmt.Println("str=", str)

	//显示控件，如果是通过glade添加的控件，show即显示所有，代码布局则需要showall
	win.Show()
	//主事件循环
	gtk.Main()
}
