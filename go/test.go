package main

import (
	"fmt"
)

func main() {
	x := make([]int, 0, 10)
	x = append(x, 1)
	y := append(x, 4, 5, 6)
	fmt.Println(x, y)
	z := append(x, 7, 8, 9)
	fmt.Println(x, y, z)
}
