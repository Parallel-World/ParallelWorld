package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
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
	//"key-press-event" 键盘按下时触发
	win.Connect("key-press-event", func(ctx *glib.CallbackContext) {
		//获取键盘按下属性结构体变量
		arg := ctx.Args(0)
		event := *(**gdk.EventKey)(unsafe.Pointer(&arg))
		//获取到的实际上是字母的ascii
		key := event.Keyval
		//fmt.Printf("key = %v\n", key)
		if key == gdk.KEY_a {
			fmt.Println("a")
		}
	})

	//通过代码添加控件，需要显示所有
	win.ShowAll()
	//主事件循环
	gtk.Main()
}
