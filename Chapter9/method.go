package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println(m);
}

func main() {
	value := MyType("a MyType value")
	value.sayHi();
}