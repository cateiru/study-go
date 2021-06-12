package main

import "fmt"

func test() (foo string) {
	foo = "Hello"
	return
}

func main() {
	a := test()
	fmt.Println(a)
}
