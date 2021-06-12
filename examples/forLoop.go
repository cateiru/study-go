package main

import (
	"fmt"
)

func main() {
	a := [3]string{"hoge", "fuga", "foo"}

	for key, value := range a {
		fmt.Println(key, value)
	}
}
