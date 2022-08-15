package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	// Root 找项目根目录
	Root = filepath.Join(filepath.Dir(b), "..")
)

func main() {
	fmt.Println("---", b)
	fmt.Println("---2", filepath.Dir(b))
	fmt.Println("---3", filepath.Join(filepath.Dir(b), ".."))
	fmt.Println("---4", Root)
	sdkWsPort := flag.Int("sdk_ws_port", 30000, "openIM ws listening port")
	fmt.Println("----5", sdkWsPort)
	fmt.Println("----5", *sdkWsPort)
}
