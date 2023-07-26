package main

import "fmt"

func calcWindow(flowerbed []int, i int) int {
	window := 0
	size := len(flowerbed)
	if i-1 >= 0 {
		window += flowerbed[i-1]
	}
	window += flowerbed[i]
	if i+1 < size {
		window += flowerbed[i+1]
	}
	return window
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	size := len(flowerbed)
	if size == 0 {
		return false
	}

	i := 0
	possibilities := 0
	for i < size {
		window := calcWindow(flowerbed, i)
		if window == 0 {
			possibilities++
			if possibilities >= n {
				return true
			}
			i += 2
			window = calcWindow(flowerbed, i)
			continue
		}
		i++
	}
	return false
}

func main() {
	flowered := []int{1, 0, 0, 0, 0, 1}
	n := 2
	fmt.Println(canPlaceFlowers(flowered, n))
}
