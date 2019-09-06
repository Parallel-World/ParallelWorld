//https://www.pengfu.com/xiaohua_1.html

package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HttpGet(url string) (resuit string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	//读取网页内容
	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		resuit += string(buf[:n]) //累加读取的内容
	}
	return
}

//爬取每个段子的title和content
func SpiderOneJoy(url string) (title, content string, err error) {
	resuit, err1 := HttpGet(url)
	if err1 != nil {
		err = err1
		return
	}
	//取标题
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>"`)
	if re1 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}
	//取内容
	tmpTitle := re1.FindAllStringSubmatch(resuit, 1) //最后一个参数为1，只过滤第一个
	for _, data := range tmpTitle {
		title = data[1]
		title = strings.Replace(title, "\t", "", -1) //将字符串中的\t换成""
		break
	}
	//取段子
	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re2 == nil {
		err = fmt.Errorf("%s", "regexp.MustCompile err2")
		return
	}
	//取内容
	tmpContent := re2.FindAllStringSubmatch(resuit, -1)
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\t", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "\r", "", -1)
		content = strings.Replace(content, "<br />", "", -1)
		break
	}
	return
}

//写入文件
func StoreJoyToFile(i int, fileTitle, fileContent []string) {
	//新建文件
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("os.Create err=", err)
		return
	}
	defer f.Close()
	//写内容
	n := len(fileTitle)
	for i := 0; i < n; i++ {
		//写标题
		f.WriteString(fileTitle[i]+"\n")
		//写内容
		f.WriteString(fileContent[i]+"\n")
		f.WriteString("\n==================")
	}
}

func SpiderPape(i int) {
	//爬取的urlhttps://www.pengfu.com/xiaohua_1.html
	url := `https://www.pengfu.com/xiaohua_` + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取%d个网页：%s\n", i, url)
	//开始爬取页面内容
	resuit, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}
	//fmt.Println("r=", resuit)		爬出整个html
	//取<h1 class="dp-b"><a href="一个段子url连接"
	//解析表达式	正则
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regexp.MustCompile err")
		return
	}
	//取关键信息
	joyUrls := re.FindAllStringSubmatch(resuit, -1)
	//fmt.Println("joyUrls = ", joyUrls)
	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)
	//取每个段子的网址
	for _, data := range joyUrls {
		//fmt.Println("url=",data[1])	data为切片，第一个为下标，第二个为内容
		//开始爬取每一个段子
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			fmt.Println("SpiderOneJoy err=", err)
			continue
		}
		//fmt.Printf("title:#%v#", title)		读取标题
		//fmt.Printf("content:#%v#", content)		读取内容
		fileTitle = append(fileTitle, title)       //追加标题内容
		fileContent = append(fileContent, content) //追加文章内容
	}
	//把内容写在文件里
	StoreJoyToFile(i, fileTitle, fileContent)
}

func DoWork(start, end int) {
	fmt.Println("准备爬取第%d页到%d页的内容", start, end)
	for i := start; i <= end; i++ {
		//定义一个函数，爬取页面
		SpiderPape(i)
	}
}

func main() {
	var start, end int
	fmt.Println("请输入起始页：")
	fmt.Scan(&start)
	fmt.Println("请输入终止页: ")
	fmt.Scan(&end)
	DoWork(start, end)
}
