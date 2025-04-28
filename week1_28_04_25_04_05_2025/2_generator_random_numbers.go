package week1_28_04_25_04_05_2025

import (
	"math/rand"
	"time"
)

func RandNumsGenerator(count int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := make(chan int)

	go func() {
		for i := 0; i < count; i++ {
			result <- r.Intn(count)
		}
		close(result)
	}()
	return result
}
