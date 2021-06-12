package main

import "fmt"

func main() {
Loop:
	for n := 0; n <= 10; n++ {
		fmt.Println("OK")

		for p := 0; p < 100; p++ {
			if p*n == 1000 {
				fmt.Println("break loop to Loop label")
				break Loop
			}
		}
	}
	fmt.Println("Loop label end")
}
