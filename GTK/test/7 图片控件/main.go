package main

import (
	"os"

	"github.com/mattn/go-gtk/gdkpixbuf"

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
	//设置图片控件
	image := gtk.ImageFromObject(builder.GetObject("image1"))
	//获取控件大小
	var w, h int
	image.GetSizeRequest(&w, &h)
	//设置图片资源，pixbuf,控件大小和图片大小一样
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("../Image/2.png", w-10, h-10, false) //false表示不保持图片原来的尺寸
	image.SetFromPixbuf(pixbuf)                                                          //给image设置图片
	//图片资源使用完毕，需要释放空间
	pixbuf.Unref()

	//显示控件，如果是通过glade添加的控件，show即显示所有，代码布局则需要showall
	win.Show()
	//主事件循环
	gtk.Main()
}
