package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64 = 0

func main() {
	// simplyFun()
	makeChan()
	// forRangeChan()
	// selectChannel()
	var w sync.WaitGroup
	w.Add(2)
	go add(&w)
	go add(&w)
	w.Wait()
	fmt.Println("count", count)
	// ch := make(chan int)
	// close(ch)
	// if _, ok := <-ch; !ok {
	// 	fmt.Println("channel 已关闭，读取不到数据")
	// }
}

func add(w *sync.WaitGroup) {
	atomic.AddInt64(&count, 1)
	w.Done()
}

// simply function
func simplyFun() {
	var wg sync.WaitGroup
	wg.Add(1)
	go run(&wg)
	wg.Wait()
}
func run(wg *sync.WaitGroup) {
	fmt.Println("go run")
	wg.Done()
}

// 定义chan
func makeChan() {
	// 有缓冲区
	intWithCache := make(chan int, 1)
	// 无缓冲区
	intWithNoCache := make(chan int)
	intWithCache <- 1
	fmt.Println("intWithCache", <-intWithCache)
	go func() {
		intWithNoCache <- 1
	}()
	fmt.Println("intWithNoCache", <-intWithNoCache)
}

// 定义chan
func forRangeChan() {
	var wg sync.WaitGroup //
	wg.Add(1)
	numWithCache := make(chan int)

	var read <-chan int = numWithCache  // 可读
	var write chan<- int = numWithCache // 可写

	go setData(write, &wg)
	go getData(read)
	wg.Wait()
}

func setData(ch chan<- int, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	wg.Done()
}

func getData(ch <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("i:", <-ch)
	}

}

func selectChannel() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})

	// ch1, ch2 send data
	go sendCh1(ch1)
	go sendCh1(ch2)

	// handle data
	for {
		select {
		case _, ok := <-ch1:
			fmt.Println("I am ch1", ok)
		case _, ok := <-ch2:
			fmt.Println("I am ch2", ok)
		}
	}
}

func sendCh1(ch chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		fmt.Println("I am send ch")
	}
	close(ch)
}

// 多个key value数据写磁盘
// Put 原本为一个put只写一个key value到磁盘
// PutMulti 执行时间和put一样，但是可以同时刷多个key，value

func Put(key, value string) error {
	return nil
}

func PutMulti(values map[string]string) error {
	return nil
}
