package main

import (
	"os"
	"strconv"

	"github.com/mattn/go-gtk/glib"

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
	//获取按钮
	buttonstart := gtk.ButtonFromObject(builder.GetObject("buttonstart"))
	buttonstop := gtk.ButtonFromObject(builder.GetObject("buttonstop"))
	//获取标签
	label := gtk.LabelFromObject(builder.GetObject("label1"))
	label.SetText("0")             //设置内容
	label.ModifyFontSize(30)       //设置字体大小
	buttonstop.SetSensitive(false) //按钮变灰
	num := 0
	id := 0
	//启动定时器
	buttonstart.Clicked(func() {
		id = glib.TimeoutAdd(500, func() bool {
			num++
			label.SetText(strconv.Itoa(num))
			return true //注意要返回true
		})
		buttonstop.SetSensitive(true)
		buttonstart.SetSensitive(false)
	})
	//停止计时器
	buttonstop.Clicked(func() {
		glib.TimeoutRemove(id)
		buttonstop.SetSensitive(false)
		buttonstart.SetSensitive(true)
	})

	//通过代码添加控件，需要显示所有
	win.ShowAll()
	//主事件循环
	gtk.Main()
}
