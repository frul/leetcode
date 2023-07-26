package main

import (
	"fmt"
)

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	this.queue = append(this.queue, t)
	count_to_delete := 0
	for _, v := range this.queue {
		if t-v > 3000 {
			count_to_delete++
		}
	}
	this.queue = this.queue[count_to_delete:]
	return len(this.queue)
}

func main() {
	fmt.Println("Hello, playground")
}
