package week1_28_04_25_04_05_2025

import "sync"

/* 3. Слить N каналов в один

Даны n каналов типа chan int. Надо написать функцию, которая смерджит все данные из этих каналов в один и вернет его.

Мы хотим, чтобы результат работы функции выглядел примерно так:

```go
for num := range joinChannels(a, b, c) {
       fmt.Println(num)
} */

func MergeChannels(channels ...chan int) chan int {
	result := make(chan int)
	wg := &sync.WaitGroup{}

	for _, channel := range channels {
		wg.Add(1)
		go func(ch <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			for num := range ch {
				result <- num
			}
		}(channel, wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}
