package main

import (
	"fmt"
	"leetcode/week1_28_04_25_04_05_2025"
)

func main() {
	for num := range week1_28_04_25_04_05_2025.RandNumsGenerator(10) {
		fmt.Println(num)
	}
}
