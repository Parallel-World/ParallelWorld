package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)
	builder := gtk.NewBuilder()
	builder.AddFromFile("test.glade")
	win := gtk.WindowFromObject(builder.GetObject("window1")) //控件名称
	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})
	//"configure_event" 窗口大小改变时触发
	win.Connect("configure_event", func() {
		var w, h int
		win.GetSize(&w, &h)
		fmt.Printf("w = %v,h=%v\n", w, h)
	})

	win.ShowAll()
	gtk.Main()
}
