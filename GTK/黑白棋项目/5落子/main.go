package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"

	"github.com/mattn/go-gtk/gdk"

	"github.com/mattn/go-gtk/gtk"
)

//控件结构体
type ChessWidget struct {
	window      *gtk.Window //窗口
	buttonMin   *gtk.Button //最小化
	buttonClose *gtk.Button //关闭
	labelBlack  *gtk.Label  //记录黑棋个数
	labelWhite  *gtk.Label  //记录白棋个数
	labelTime   *gtk.Label  //记录倒计时
	imageBlack  *gtk.Image  //提示黑子落子
	imageWhite  *gtk.Image  //提示白字落子
}

//控件属性结构体
type ChessInfo struct {
	w, h           int //窗口的宽度和高度
	x, y           int //鼠标点击窗口的坐标
	startX, startY int //棋盘起点坐标
	gridW, gridH   int //棋盘一个格子的宽度和高度

}

const (
	Empty = iota //当前棋盘格子没有子
	Black        //当前棋盘格子为黑子
	White        //当前棋盘格子为白子
)

//黑白棋结构体
type Chessboard struct {
	ChessWidget //匿名字段
	ChessInfo

	currentRole int //该谁落子
	tipTimerid  int //实现提示闪烁效果的定时器id

	chess [8][8]int //二维数组标记棋盘

}

//函数：给按钮设置图片
func ButtonSetImageFromFile(button *gtk.Button, filename string) {
	//获取按钮大小
	w, h := 0, 0
	button.GetSizeRequest(&w, &h)
	//创建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, w-10, h-10, false)
	//创建image
	image := gtk.NewImageFromPixbuf(pixbuf)
	//释放pixbuf
	pixbuf.Unref()
	//给按钮设置图片
	button.SetImage(image)
	//去掉按钮焦距
	button.SetCanFocus(false)
}

//给image设置图片
func ImageSetPicFromFile(image *gtk.Image, filename string) {
	//获取image的大小
	w, h := 0, 0
	image.GetSizeRequest(&w, &h)
	//创建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, w-10, h-10, false)
	//给image设置图片
	image.SetFromPixbuf(pixbuf)
	//释放pixbuf
	pixbuf.Unref()
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
	//按钮相关
	//获取按钮控件
	obj.buttonMin = gtk.ButtonFromObject(builder.GetObject("buttonMin"))
	obj.buttonClose = gtk.ButtonFromObject(builder.GetObject("buttonClose"))
	//给按钮插图标
	ButtonSetImageFromFile(obj.buttonMin, "../Image/Min.png")
	ButtonSetImageFromFile(obj.buttonClose, "../Image/Close.png")
	//标签相关
	obj.labelBlack = gtk.LabelFromObject(builder.GetObject("labelBlack"))
	obj.labelWhite = gtk.LabelFromObject(builder.GetObject("labelWhite"))
	obj.labelTime = gtk.LabelFromObject(builder.GetObject("labelTime"))
	//设置字体大小
	obj.labelBlack.ModifyFontSize(50)
	obj.labelWhite.ModifyFontSize(50)
	obj.labelTime.ModifyFontSize(30)
	//设置内容
	obj.labelBlack.SetText("2")
	obj.labelWhite.SetText("2")
	obj.labelTime.SetText("20")
	//改变字体颜色
	obj.labelBlack.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	obj.labelWhite.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	obj.labelTime.ModifyFG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	//image相关
	obj.imageBlack = gtk.ImageFromObject(builder.GetObject("imageBlack"))
	obj.imageWhite = gtk.ImageFromObject(builder.GetObject("imageWhite"))
	//设置图标
	ImageSetPicFromFile(obj.imageBlack, "../Image/alex.png")
	ImageSetPicFromFile(obj.imageWhite, "../Image/peiqi.png")
	//棋盘坐标相关
	obj.startX, obj.startY = 200, 60
	obj.gridW, obj.gridH = 50, 40
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
	i := (obj.x - obj.startX) / obj.gridW
	j := (obj.x - obj.startY) / obj.gridH
	if i >= 0 && i <= 7 && j >= 0 && j <= 7 {
		fmt.Printf("(%d,%d)\n", i, j)
		obj.chess[i][j] = Black
		//刷新绘图区域(整个窗口)
		obj.window.QueueDraw()
	}

}

//鼠标移动事件
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

func PaintEvent(ctx *glib.CallbackContext) {
	//获取用户传递参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard) //类型断言
	if ok == false {
		fmt.Println("MouseMoveEcent Chessboard err")
		return
	}
	//获取画家，设置绘图区域
	painter := obj.window.GetWindow().GetDrawable()
	gc := gdk.NewGC(painter)
	//新建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("../Image/1.png", obj.w, obj.h, false)
	//黑白棋pixbuf
	blackPixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("../Image/A.png", obj.gridW, obj.gridH, false)
	whitePixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("../Image/B.png", obj.gridW, obj.gridH, false)

	//画图
	painter.DrawPixbuf(gc, pixbuf, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
	//画黑白棋
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if obj.chess[i][j] == Black {
				fmt.Println(obj.chess[i][j])
				painter.DrawPixbuf(gc, blackPixbuf, 0, 0, obj.startX+i*obj.gridW, obj.startY+j*obj.gridH, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			} else if obj.chess[i][j] == White {
				fmt.Println(obj.chess[i][j])
				painter.DrawPixbuf(gc, whitePixbuf, 0, 0, obj.startX+i*obj.gridW, obj.startY+j*obj.gridH, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
			}
		}
	}
	//释放资源
	pixbuf.Unref()
	blackPixbuf.Unref()
	whitePixbuf.Unref()
}

//方法：事件、信号处理
func (obj *Chessboard) HandleSignal() {
	//鼠标点击事件
	//"button-press-event":鼠标按下时触发
	obj.window.Connect("button-press-event", MousePressEvent, obj)
	//鼠标移动事件
	//"motion-notify-event"按住鼠标移动时触发
	obj.window.Connect("motion-notify-event", MouseMoveEvent, obj)
	//按钮的信号处理
	obj.buttonClose.Clicked(func() {
		//关闭定时器
		glib.TimeoutRemove(obj.tipTimerid)
		gtk.MainQuit() //关闭窗口
	})
	obj.buttonMin.Clicked(func() {
		obj.window.Iconify() //最小化窗口
	})
	//绘图相关
	//大小改变事件
	//"configure_event" 窗口大小改变时触发
	obj.window.Connect("configure_event", func() {
		//重新刷图
		obj.window.QueueDraw()
	})
	//绘图事件，"expose-event"
	obj.window.Connect("expose-event", PaintEvent, obj)

}

//函数：提示功能，实现闪烁效果
func ShowTip(obj *Chessboard) {
	if obj.currentRole == Black { //当前黑子下
		//隐藏白子
		obj.imageWhite.Hide()
		if obj.imageBlack.GetVisible() == true {
			//原来是显示的需要隐藏
			obj.imageBlack.Hide()
		} else { //原来是隐藏的需要显示
			obj.imageBlack.Show()
		}
	} else { //当前白子下
		//隐藏黑子
		obj.imageBlack.Hide()
		if obj.imageWhite.GetVisible() == true {
			//原来是显示的需要隐藏
			obj.imageWhite.Hide()
		} else { //原来是隐藏的需要显示
			obj.imageWhite.Show()
		}
	}
}

//方法：黑白棋属性相关
func (obj *Chessboard) InitChess() {
	//image都先隐藏
	obj.imageBlack.Hide()
	obj.imageWhite.Hide()
	//默认黑子先下
	obj.currentRole = Black

	//启动定时器
	obj.tipTimerid = glib.TimeoutAdd(500, func() bool {
		ShowTip(obj)
		return true
	})
}

func main() {
	gtk.Init(&os.Args)
	var obj Chessboard //创建结构体变量
	obj.CreateWindow() //创建控件，设置控件属性
	obj.HandleSignal() //事件信号处理
	obj.InitChess()
	obj.window.Show() //显示控件

	gtk.Main()
}
