package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"studyGo/网络编程/proto"
)

// func process(conn net.Conn) {
// 	defer conn.Close()
// 	for {
// 		var buf [128]byte
// 		n, err := conn.Read(buf[:])
// 		if err != nil {
// 			fmt.Printf("read from conn failed,err:%v\n", err)
// 			break
// 		}
// 		str := string(buf[:n])
// 		fmt.Printf("recv from lient,data:%v\n", str)
// 	}
// }
// func main() {
// 	listen, err := net.Listen("tcp", "0.0.0.0:8080")
// 	if err != nil {
// 		fmt.Println("listen failed,err:", err)
// 		return
// 	}
// 	for {
// 		conn, err := listen.Accept()
// 		if err != nil {
// 			fmt.Printf("accept failed,err:%v\n", err)
// 			continue
// 		}
// 		go process(conn)
// 	}
// }

// func process(conn net.Conn) {
// 	defer conn.Close()
// 	reader := bufio.NewReader(conn)
// 	var buf [1024]byte
// 	for {
// 		n, err := reader.Read(buf[:])
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("read from client failed,err:", err)
// 			break
// 		}
// 		recvStr := string(buf[:n])
// 		fmt.Println("收到clinet发来的数据:", recvStr)
// 	}
// }
// func main() {
// 	listen, err := net.Listen("tcp", "127.0.0.1:8080")
// 	if err != nil {
// 		fmt.Println("listen failed,err:", err)
// 		return
// 	}
// 	defer listen.Close()
// 	for {
// 		conn, err := listen.Accept()
// 		if err != nil {
// 			fmt.Println("accept failed,err:", err)
// 			continue
// 		}
// 		go process(conn)
// 	}
// }

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed,err:", err)
			return
		}
		fmt.Println("收到client发来的数据:", msg)
	}
}
func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}
		go process(conn)
	}
}
