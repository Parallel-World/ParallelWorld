package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/glib"

	"github.com/mattn/go-gtk/gdk"

	"github.com/mattn/go-gtk/gtk"
)

//控件结构体
type ChessWidget struct {
	window *gtk.Window
}

//控件属性结构体
type ChessInfo struct {
	w, h int //窗口的宽度和高度
	x, y int //鼠标点击窗口的坐标
}

//黑白棋结构体
type Chessboard struct {
	ChessWidget //匿名字段
	ChessInfo
}

func (obj *Chessboard) CreateWindow() {
	//读取glade
	builder := gtk.NewBuilder()
	builder.AddFromFile("ui.glade")
	//窗口相关
	obj.window = gtk.WindowFromObject(builder.GetObject("window1")) //获取控件
	obj.window.SetAppPaintable(true)                                //允许绘图
	obj.window.SetPosition(gtk.WIN_POS_CENTER)                      //居中显示
	obj.w, obj.h = 800, 480                                         //窗口的宽度和高度
	obj.window.SetSizeRequest(800, 480)                             //设置窗口的宽高
	obj.window.SetDecorated(false)                                  //去除边框
	//设置事件，让窗口可以捕获鼠标点击和移动
	obj.window.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))
}

//鼠标点击事件函数
func MousePressEvent(ctx *glib.CallbackContext) {
	//获取用户传递的参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard) //类型断言
	if ok == false {
		fmt.Println("MousePressEvent Chessboard err")
		return
	}
	//获取鼠标按下属性结构体变量
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
	//保存点击的x,y坐标
	obj.x, obj.y = int(event.X), int(event.Y)
	fmt.Println("x=", obj.x, ",y=", obj.y)
}

func MouseMoveEvent(ctx *glib.CallbackContext) {
	//获取用户传递的参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard) //类型断言
	if ok == false {
		fmt.Println("MouseMoveEvent Chessboard err")
		return
	}
	//获取鼠标按下属性结构体变量
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
	//保存点击的x,y坐标
	x, y := int(event.XRoot)-obj.x, int(event.YRoot)-obj.y
	obj.window.Move(x, y) //窗口移动
}

//方法：事件、信号处理
func (obj *Chessboard) HandleSignal() {
	//鼠标点击事件
	//"button-press-event":鼠标按下时触发
	obj.window.Connect("button-press-event", MousePressEvent, obj)
	//鼠标移动事件
	//"motion-notify-event"按住鼠标移动时触发
	obj.window.Connect("button-notify-event", MouseMoveEvent, obj)
}

func main() {
	gtk.Init(&os.Args)
	var obj Chessboard //创建结构体变量
	obj.CreateWindow() //创建控件，设置控件属性
	obj.HandleSignal() //事件信号处理

	obj.window.Show() //显示控件

	gtk.Main()
}
