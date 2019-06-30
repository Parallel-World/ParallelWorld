package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// func test() {
// 	content, err := ioutil.ReadFile("./file.go")
// 	if err != nil {
// 		fmt.Println("read file failed,err:", err)
// 		return
// 	}
// 	fmt.Println(string(content))
// }

// func test() {
// 	inputFile, err := os.Open("文件名.gz") //只读的方式打开
// 	if err != nil {
// 		fmt.Printf("open file err:%v\n", err)
// 		return
// 	}
// 	defer inputFile.Close()
// 	reader, err := gzip.NewReader(file)
// 	if err != nil {
// 		fmt.Println("gzip new reader failed,err", err)
// 		return
// 	}
// 	var content []byte
// 	var buf [128]byte
// 	for {
// 		n, err := reader.reader.Read(buf[:])
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("read file:", err)
// 			return
// 		}
// 		content = append(content, buf[:n]...)
// 	}
// 	fmt.Println(string(content))
// }

// func test() {
// file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
// if err != nil {
// 	fmt.Println("open file failed,err:", err)
// 	return
// }
// defer file.Close()
// 	str := "hello world"
// 	file.Write([]byte(str))
// 	file.WriteString(str)
// }

// func test() {
// 	file, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
// 	if err != nil {
// 		fmt.Println("open file failed,err:", err)
// 		return
// 	}
// 	defer file.Close()
// 	writer := bufio.NewWriter(file)
// 	for i := 0; i < 10; i++ {
// 		writer.WriteString("hello world")
// 	}
// 	writer.Flush()
// }

// func test(){
// 	str :="hello world"
// 	err := ioutil.WriteFile("./test.txt",[]byte(str),0755)
// 	if err !=nil{
// 		fmt.Println("write file failed,err:",err)
// 		return
// 	}
// }

// func test() {
// 	_,err :=CopyFile("target.txt", "source.txt")
// 	if err !=nil{
// 		fmt.Printf("copy file failed,err:%v\n",err)
// 		return
// 	}
// 	fmt.Println("Copy done!")
// }
// func CopyFile(dstName, srcName string) (written int64, err error) {
// 	src, err := os.Open(srcName)
// 	if err != nil {
// 		fmt.Printf("open source file %s failed,err:%v\n",srcName,err)
// 		return
// 	}
// 	defer src.Close()
// 	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
// 	if err != nil {
// 		fmt.Printf("open dest file %s failed,err:%v\n",dstName,err)
// 		return
// 	}
// 	defer dst.Close()
// 	return io.Copy(dst, src)
// }

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
		return
	}
}
func test() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n",
				os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}

func main() {
	test()
}
