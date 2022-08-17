package main

import (
	"fmt"
	"reflect"
)

func main() {
	animal := Animal{}
	value := reflect.ValueOf(&animal)
	f := value.MethodByName("Eat")
	call := f.Call([]reflect.Value{})
	fmt.Println("cal:", call)
}

type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("I am Eating")
}
