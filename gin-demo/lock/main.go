package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64

func add(w *sync.WaitGroup) {
	atomic.AddInt64(&count, 1)
	w.Done()
}

var flag int64

func addCAS(w *sync.WaitGroup) {
	for {
		// 此处如果无法进行交换，则代表数值不为0, 原子交换
		// 可以使用此方式进行同步操作
		if atomic.CompareAndSwapInt64(&flag, 0, 1) {
			count++
			// 原子存储，讲0存到flag地址中
			atomic.StoreInt64(&flag, 0)
			w.Done()
			return
		}
	}
}

var m sync.Mutex

func addMutex() {
	m.Lock()
	count++
	m.Unlock()
}

// 读写锁
type stat struct {
	counters map[string]int64
	mutex    sync.RWMutex
}

func (s *stat) getCounter(name string) int64 {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return s.counters[name]
}
func (s *stat) setCounter(name string) {
	s.mutex.Lock()
	defer s.mutex.Lock()
	s.counters[name]++
}
func main() {
	var w sync.WaitGroup
	w.Add(100)
	// go add(&w)
	// go add(&w)
	for i := 0; i < 100; i++ {
		go addCAS(&w)
	}
	w.Wait()
	fmt.Println("count", count)
}
