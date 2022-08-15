package main

import (
	"fmt"
	"time"
)

var intChan chan int

func isPrimeA(intChan chan int, primeChan chan int, exitChan chan bool) {
	var flag bool
label:
	for {
		select {
		case num := <-intChan:
			flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				primeChan <- num
			}
		default:
			break label
		}

	}
	fmt.Println("协程已结束")
	exitChan <- true
}

func initChan(num int) {
	for i := 1; i <= num; i++ {
		intChan <- i
	}
}

func main() {
	start := time.Now()
	intChan = make(chan int, 100)
	go initChan(100)
	var primeChan chan int = make(chan int, 100)
	var exitChan chan bool = make(chan bool, 8)
	// 有缓冲区，如果在先写后读的情况下，会先写入buf，读的goroutine在读的时候先读buf，并且会将send队列的数据移到buf。
	// 有缓冲区，如果先读后写，则会挂载sendq队列，然后调起休眠gopark，等到有goroutine来读取数据了，则会从刚刚的goroutine把数据copy过来；
	// 并对其设置goready，这样读的goroutine就会移除等待队列，等待下次调度。
	for i := 0; i < 8; i++ {
		go isPrimeA(intChan, primeChan, exitChan)
	}
	go func() {
		for i := 0; i < 7; i++ {
			<-exitChan
		}
		// close(primeChan)
	}()
	// 读取素数
	go func() {
	label:
		for {
			select {
			// select无论设计channel操作是否阻塞，都不会被阻塞（此时for range 读书channel，如果是先读后写的情况，是会在recvq队列中阻塞等待。）
			case res, ok := <-primeChan:
				if !ok {
					fmt.Println("out")
					break label
				}
				fmt.Println("素数：", res)
				// default:
				// 	break label
			}
		}
	}()
	end := time.Since(start)
	fmt.Println("用时：", end) // 100000  824.848ms
}
