package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	// 存储生产者生产产品管道
	storageChan := make(chan Product, 1000)
	// 存储到店转运管道
	shopChan := make(chan Product, 1000)
	exitChan := make(chan bool, 1000)
	for i := 1; i < 999; i++ {
		go Producer(storageChan, 1000)
	}
	go Logistics(storageChan, shopChan)
	go Consumer(shopChan, 1000*999, exitChan)
	if <-exitChan {
		return
	}
}

// 商品
type Product struct {
	Name string
}

var wl sync.Mutex

// 生产者
func Producer(storageChan chan<- Product, count int) {
	tempCount := count
	for {
		producer := Product{"商品：" + strconv.Itoa(tempCount)}
		storageChan <- producer
		tempCount--
		time.Sleep(time.Second)
		// fmt.Println("生产了", producer)
		if tempCount < 1 {
			return
		}
	}
}

// 物流公司
func Logistics(storageChan <-chan Product, shopChan chan<- Product) {
	for {
		product := <-storageChan
		shopChan <- product
		// fmt.Println("运输了", product)
	}
}

// 消费者
func Consumer(shopChan <-chan Product, count int, exitChan chan<- bool) {
	for {
		product := <-shopChan
		fmt.Println("消费了", product)
		count--
		if count < 1 {
			exitChan <- true
			break
		}
	}
}
