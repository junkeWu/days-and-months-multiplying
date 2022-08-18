package main

import (
	"fmt"
	"sync"
)

func main() {
	var sm sync.Map

	sm.Store("Hello", "hello world!")
	sm.Store("k2", "v2")
	fmt.Println(sm.Load("k2"))

	sm.Range(func(key, value interface{}) bool {
		fmt.Println("iterate:", key, value)
		return true
	})
}
