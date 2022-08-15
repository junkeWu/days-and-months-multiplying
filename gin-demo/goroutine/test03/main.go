package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*	// 方式一：
		fmt.Println("当前时间：", time.Now())
		// timer := time.NewTimer(time.Second * 3)
		// t := <-timer.C // timer.C 是一个只读的管道

		// fmt.Println(t)
		// 方式二
		t := <-time.After(time.Second * 3) // 源码可见 实际上返回的就是   return NewTimer(d).C
		fmt.Println(t)
	*/

	var count = 0
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		defer waitGroup.Done()
		defer ticker.Stop()
		for {
			t := <-ticker.C
			count++
			if count > 4 {
				return
			}
			fmt.Println("时间：", t.Format("2006-01-02 03:04:05"))
		}
	}()
	// time.Sleep(time.Second * 10)
	waitGroup.Wait()
	fmt.Println("游戏结束") //

}
