package main

import "fmt"

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

func add() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)

	fmt.Println(x)
}

func main() {
	arr := [...]float64{7.0, 8.5, 9.1}
	x := Sum(&arr)

	fmt.Println(x)

	add()
}
