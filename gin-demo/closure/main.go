package main

import "fmt"

func main() {
	f := add()
	for i := 0; i < 10; i++ {
		fmt.Println("add:", f())
	}

	fi := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fi())
	}

	fmt.Println(fn())

}

func add() func() int {
	sum := 0
	return func() int {
		sum += 1
		return sum
	}
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type Cursor struct {
	X int
}

// go build --gcflags=-m main.go
// 生成汇编 go tool compile -S main.go
// go: noinline
func fn() *Cursor {
	var c Cursor
	c.X = 500
	return &c
}
