package main

import "fmt"

type myConst int

const (
	zero  myConst = iota // iota = 0
	one                  // iota = 1
	two                  // iota = 2
	three                // iota = 2
	foure                // iota = 3
	five                 // iota = 4
)
const (
	EOF = -(iota + 2)
	Ident
	Int
	Float
	Char
	String
	RawString
	Comment

	// internal use only
	skipComment
)

func main() {
	// 3 4 5 why not 3 4 6
	fmt.Println(zero, one, two, three, foure, five)
	fmt.Println(EOF, Ident, Int, Float, Char, String, RawString)
}
