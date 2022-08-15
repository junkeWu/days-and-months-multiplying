package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

func main2() {
	resp, err := http.Get("http://119.23.246.18:8080")
	if err != nil {
		errors.New("error")
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	bytes := make([]byte, 1024)
	resp.Body.Read(bytes)
	if err != nil && err != io.EOF {
		fmt.Println(err)
		return
	} else {
		fmt.Println("all over are read")
		res := string(bytes[:])
		fmt.Println(res)
	}
}
func main() {
	dial, err := net.Dial("tcp", "119.23.246.18:8080")
	if err != nil {
		return
	}
	defer dial.Close()
	write, err := dial.Write([]byte("hello, I am client"))
	fmt.Println("err", err)
	fmt.Println("write", write)
}
