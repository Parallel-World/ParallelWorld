package main

import (
	"os"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"

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
	var w, h int
	win.Connect("configure_event", func() {
		win.GetSize(&w, &h)
		//刷图，这个窗口区域刷图
		win.QueueDraw() //触发"expose-event"
	})
	//允许窗口绘图
	win.SetAppPaintable(true)
	x := 0
	y := w
	//绘图是触发的信号：expose-event
	win.Connect("expose-event", func() {
		//设置画家，指定绘图区域
		painter := win.GetWindow().GetDrawable()
		gc := gdk.NewGC(painter)
		//创建图片资源
		bg, _ := gdkpixbuf.NewPixbufFromFileAtScale("../Image/6.png", w, h, false)
		face, _ := gdkpixbuf.NewPixbufFromFileAtScale("../Image/3.png", 100, 100, false)
		face3, _ := gdkpixbuf.NewPixbufFromFileAtScale("../Image/7.png", 100, 100, false)

		painter.DrawPixbuf(gc, bg, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
		painter.DrawPixbuf(gc, face, 0, 0, x, 150, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
		painter.DrawPixbuf(gc, face3, 0, 0, x, 300, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

		//释放图片资源
		bg.Unref()
		face.Unref()
		face3.Unref()
	})
	//获取按钮
	button := gtk.ButtonFromObject(builder.GetObject("button1"))
	button.Clicked(func() {
		x += 120
		if x >= w {
			x = 0
		}
		y -= 80
		if y <= 0 {
			y = w
		}
		//刷图，这个窗口区域刷图
		win.QueueDraw()
	})

	win.ShowAll()
	gtk.Main()
}
