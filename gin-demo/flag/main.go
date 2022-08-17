package main

import (
	"flag"
	"fmt"
)

// go run main.go -port=8080 可以运行用参数改变port
func main() {
	port := flag.Int("port", 8000, "rpc listen port")
	flag.Parse()
	fmt.Println("port", *port)

	fmt.Println("-----------------------")
	flag.Parse() // 解析参数
	fmt.Printf("%s:%d\n", host, portG)
}

var host string
var portG int

func init() { // 每个文件会自动执行的函数
	flag.StringVar(&host, "host", "127.0.0.1", "请输入host地址")
	flag.IntVar(&portG, "portG", 3306, "请输入端口号")
}
