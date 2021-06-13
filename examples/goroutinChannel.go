package main

import "fmt"

func caseOne() {
	message := make(chan string)

	go func() {
		message <- "Hello"
	}()

	msg := <-message // wait

	fmt.Println(msg)
}

func caseTwo() {
	var sem = make(chan int, MaxOutStanging)
}

func main() {
	// caseOne()
}
