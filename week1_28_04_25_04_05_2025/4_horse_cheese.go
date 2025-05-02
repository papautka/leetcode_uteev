package week1_28_04_25_04_05_2025

import "fmt"

func CountCombination2(N int, M int) int {
	// 1. Инициализируем массив нулями
	dp := make([][]int, N)
	for i := range dp {
		dp[i] = make([]int, M)
	}

	// 2. Задаем стартовую точку (левый верхний угол)
	dp[0][0] = 1

	// 2.1 Проверяем можно ли походить с позиции
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			// 2.1.1 (2 вниз 1 вправо)
			if i-2 >= 0 && j-1 >= 0 {
				dp[i][j] = dp[i-2][j-1] + dp[i-2][j-1]
			}
			// 2.1.2 (1 вниз 2 вправо)
			if i-1 >= 0 && j-2 >= 0 {
				dp[i][j] = dp[i-1][j-2] + dp[i-1][j-2]
			}
		}
	}

	// 3. Для наглядности печатаем результат
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			fmt.Printf("%d ", dp[i][j])
		}
		fmt.Printf("\n")
	}
	return dp[N-1][M-1]

}
