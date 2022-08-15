package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

/*
func main() {
	// http://127.0.0.1:8080/go
	// 单独写回调
	http.HandleFunc("/go", myHandler)
	// addr: 监听的地址
	// handler: 回调函数
	http.ListenAndServe("127.0.0.1:8080", nil)
}
*/

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}
	defer listen.Close()
	time.Sleep(time.Second * 10)
	for {
		// 从socket recive队列里获取一个建立好的连接
		time.Sleep(time.Second * 10)
		conn, err := listen.Accept()
		fmt.Println("非阻塞")
		if err != nil {
			return
		}
		// 新起一个goroutine处理连接
		go Handler(conn)
	}

}

func Handler(conn net.Conn) {
	defer conn.Close()
	var dataBuffer bytes.Buffer
	buf := make([]byte, 1024)
	for {

		n, err := conn.Read(buf) // 从conn中读取客户端发送的数据内容
		if err != nil {
			if err == io.EOF {
				fmt.Printf("客户端退出 err=%v\n", err)
			} else {
				fmt.Printf("read err=%v\n", err)
			}
			break
		}
		dataBuffer.Write(buf[:n])
	}

	fmt.Println("server:", dataBuffer.String())
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RemoteAddr, "connetion success")
	// 请求方式
	fmt.Println("method", r.Method)
	// go
	fmt.Println("url", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body:", r.Body)
	//
	w.Write([]byte("www.com"))
	fmt.Println(json.Marshal(r.URL))
}
