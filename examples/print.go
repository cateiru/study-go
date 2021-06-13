package main

import "fmt"

func mapPrint() {
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	fmt.Println(timeZone)         // map[CST:-21600 EST:-18000 MST:-25200 PST:-28800 UTC:0]
	fmt.Printf("%#v\n", timeZone) // map[string]int{"CST":-21600, "EST":-18000, "MST":-25200, "PST":-28800, "UTC":0}
	fmt.Printf("%T\n", timeZone)  // map[string]int
}

func main() {
	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
	mapPrint()
}
