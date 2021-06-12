package main

import "fmt"

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func main() {
	arr := [...]float64{7.0, 8.5, 9.1}
	x := Sum(&arr)

	fmt.Println(x)
}
