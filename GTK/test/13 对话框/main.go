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
	//获取按钮
	b1 := gtk.ButtonFromObject(builder.GetObject("button1"))
	b2 := gtk.ButtonFromObject(builder.GetObject("button2"))
	//问题对话框
	b1.Clicked(func() {
		dialog := gtk.NewMessageDialog(
			win,                  //指定父窗口
			gtk.DIALOG_MODAL,     //模态对话框,对话框不关闭，父窗口不能移动
			gtk.MESSAGE_QUESTION, //问题对话框
			gtk.BUTTONS_YES_NO,   //按钮
			"这是问题对话框")            //这是内容
		dialog.SetTitle("消息问题框")
		//运行，然后销毁
		ret := dialog.Run()
		if ret == gtk.RESPONSE_YES {
			fmt.Println("yes")
		} else if ret == gtk.RESPONSE_NO {
			fmt.Println("no")
		} else {
			fmt.Println("close")
		}
		dialog.Destroy()
	})
	//消息对话框
	b2.Clicked(func() {
		dialog := gtk.NewMessageDialog(
			win,              //指定父窗口
			gtk.DIALOG_MODAL, //模态对话框,对话框不关闭，父窗口不能移动
			gtk.MESSAGE_INFO, //消息对话框
			gtk.BUTTONS_OK,   //按钮
			"这是消息对话框")        //这是内容
		dialog.SetTitle("消息对话框")
		//运行，然后销毁
		dialog.Run()
		dialog.Destroy()
	})

	//通过代码添加控件，需要显示所有
	win.ShowAll()
	//主事件循环
	gtk.Main()
}
