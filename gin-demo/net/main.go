package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listen, _ := net.Listen("tcp", "127.0.0.1:8000")
	fmt.Println("listen", listen)

	time.Sleep(time.Second * 10)
}
