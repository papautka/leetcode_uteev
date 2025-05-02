package week1_28_04_25_04_05_2025

import (
	"sync"
)

// a(ch) [1, 2, 3]
// b(ch) [4, 5]
// c(ch) [6, 7, 8, 9]

// res(ch) [1, 2, 3, 4, 5, 6, 7, 8, 9 ]

type Collected struct {
	data []int
	idx  int
}

func MergeChannelStrong(chs ...chan int) chan int {
	result := make(chan int)

	go func() {
		defer close(result)

		var wg sync.WaitGroup
		collector := make(chan Collected, len(chs))

		// собираем все каналы параллельно
		for idx, ch := range chs {
			wg.Add(1)
			go func(idx int, ch chan int) {
				defer wg.Done()
				var slice []int
				for val := range ch {
					slice = append(slice, val)
				}
				collector <- Collected{slice, idx}
			}(idx, ch)
		}

		// отдельная горутина для закрытия collector
		go func() {
			wg.Wait()
			close(collector)
		}()

		// собираем все в map
		grouped := make(map[int][]int)
		for v := range collector {
			grouped[v.idx] = v.data
		}

		// выводим в правильном порядке
		for i := 0; i < len(chs); i++ {
			for _, v := range grouped[i] {
				result <- v
			}
		}
	}()

	return result
}

/*
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

	for elemCh := range week1_28_04_25_04_05_2025.MergeChannelStrong(a, b, c) {
		fmt.Println(elemCh)
	}
}

*/
