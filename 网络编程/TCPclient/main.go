package main

import (
	"fmt"
	"net"
	"studyGo/网络编程/proto"
)

// func main() {
// 	conn, err := net.Dial("tcp", "0.0.0.0:8080")
// 	if err != nil {
// 		fmt.Println("dial failed,err:%v\n", err)
// 		return
// 	}
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		data, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Printf("read from console faild,err:%v\n", err)
// 			break
// 		}
// 		data = strings.TrimSpace(data)
// 		if data == "Q" {
// 			return
// 		}
// 		_, err = conn.Write([]byte(data))
// 		if err != nil {
// 			fmt.Printf("write failed,err:%v\n", err)
// 			break
// 		}
// 	}
// }

// func main() {
// 	conn, err := net.Dial("tcp", "www.baidu.com:80")
// 	if err != nil {
// 		fmt.Printf("dial failed,err:%v\n", err)
// 		return
// 	}
// 	data := "GET / HTTP/1.1\r\n"
// 	data += "HOST:www.baidu.com\r\n"
// 	data += "connection: close\r\n"
// 	data += "\r\n\r\n"
// 	_, err = io.WriteString(conn, data)
// 	if err != nil {
// 		fmt.Printf("write string failed,err:%v\n", err)
// 		return
// 	}
// 	var buf [1024]byte
// 	for {
// 		n, err := conn.Read(buf[:])
// 		if err != nil || n == 0 {
// 			break
// 		}
// 		fmt.Println(string(buf[:n]))
// 	}
// }

// func main() {
// 	conn, err := net.Dial("tcp", "127.0.0.1:8080")
// 	if err != nil {
// 		fmt.Println("dial failed,err:", err)
// 		return
// 	}
// 	defer conn.Close()
// 	for i := 0; i < 20; i++ {
// 		msg := `Hello,Hello2.How are you?`
// 		conn.Write([]byte(msg))
// 	}
// }

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("dial failed,err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello,Hello.How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed,err", err)
			return
		}
		conn.Write(data)
	}
}
