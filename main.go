package main

import (
	"fmt"
	"leetcode/week1_28_04_25_04_05_2025"
)

func main() {

	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, val := range []int{1, 2} {
			a <- val
		}
		close(a)
	}()

	go func() {
		for _, val := range []int{3, 4, 5} {
			b <- val
		}
		close(b)
	}()

	go func() {
		for val := range []int{6, 7, 8, 9, 10} {
			c <- val
		}
		close(c)
	}()

	for num := range week1_28_04_25_04_05_2025.MergeChannels(a, b, c) {
		fmt.Println(num)
	}
}
