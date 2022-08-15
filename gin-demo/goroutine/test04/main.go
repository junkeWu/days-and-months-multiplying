package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	cancel, cancelFunc := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(cancel)
	time.Sleep(time.Second * 2)
	cancelFunc()
	wg.Wait()
	fmt.Println("over")
}

func worker(ctx context.Context) {
NEXT:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			break NEXT
		default:
		}
	}
	wg.Done()
}
