package main

import "sync"

func main() {
	once := sync.Once{}

	once.Do(func() {

	})
}
