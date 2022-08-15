package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64 = 0

func add(w *sync.WaitGroup) {
	atomic.AddInt64(&count, 1)
	w.Done()
}

func main() {

	slice01 := []int{1, 2, 3}
	slice02 := append(slice01, []int{4, 5, 6}...)
	fmt.Println("cap:", cap(slice01), len(slice01), slice01)
	fmt.Println("cap:", cap(slice02), len(slice02), slice02) // 1 2 4 8 16 32 64
	var w sync.WaitGroup
	w.Add(2)
	go add(&w)
	go add(&w)
	w.Wait()
	fmt.Println("count", count)

	intChan := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan <- i * 2 // 放入100个数据到管道
	}

	// 遍历管道时，不能使用普通的for循环结构
	// for i := 0; i < len(intChan); i++ {
	// }

	// 如果在遍历时没有关闭channel，则会出现deadlock错误
	// close(intChan)
	// for v := range intChan {
	// 	fmt.Println("v=", v)
	// }

}
