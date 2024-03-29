package main

import (
	"fmt"
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
	//获取按钮控件
	b1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	b2 := gtk.ButtonFromObject(builder.GetObject("button2"))
	b1.SetLabel("^_^") //设置文本内容
	b1.SetLabelFontSize(30)
	fmt.Println("b1 text = ", b1.GetLabel()) //获取内容
	b1.SetSensitive(false)                   //变灰，不让按             //变灰，不让按
	//获取b2大小
	var w, h int
	b2.GetSizeRequest(&w, &h)
	//新建图片资源，大小和b2差不多
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("../Image/1.png", w-10, h-10, false)
	//新建image
	image := gtk.NewImageFromPixbuf(pixbuf)
	pixbuf.Unref()        //释放资源
	b2.SetImage(image)    //按钮设置图标
	b2.SetCanFocus(false) //取消边框

	//显示控件，如果是通过glade添加的控件，show即显示所有，代码布局则需要showall
	win.Show()
	//主事件循环
	gtk.Main()
}
