package main

func sum(list []int, sum_channel chan int) {
	var sum int
	for _, v := range list {
		sum += v
	}
	sum_channel <- sum
}

func main() {
	var list []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	var sum_channel chan int = make(chan int)
	go sum(list[:len(list)/2], sum_channel)
	go sum(list[len(list)/2:], sum_channel)
	var x, y int = <-sum_channel, <-sum_channel
	println(x + y)
}
