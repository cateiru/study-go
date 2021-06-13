package main

import (
	"fmt"
	"time"
)

func printer(text string) {
	for i := 0; 100 > i; i++ {
		fmt.Printf("%s: %d\n", text, i)
		time.Sleep(100 * time.Microsecond)
	}

	fmt.Printf("End %s\n", text)
}

func main() {
	go printer("Hello")
	printer("world")
}
