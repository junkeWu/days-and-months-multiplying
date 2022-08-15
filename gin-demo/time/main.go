package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("I am a main.go")
	timer := time.NewTimer(time.Second)
EXIT:
	for {
		fmt.Println("=====第一层")
		select {
		case <-timer.C: // 只会触发一次，定时器，
			fmt.Println("timeout")
		case <-time.After(time.Millisecond * time.Duration(5*1000)): // 计时  == time.second*5
			fmt.Println("1秒钟")
			break EXIT
		}

		fmt.Println("over")
	}
}
