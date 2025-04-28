package week1_28_04_25_04_05_2025

// a := []int{23, 3, 1, 2}
// b := []int{6, 2, 4, 23}

func Intersertion(a, b []int) []int {
	// 1. Создаем map
	counter := make(map[int]int)

	// 2. Пустой слайс для result
	result := make([]int, 0)

	// 3. Результатом этого цикла является [1:1 2:1 3:1 23:1]
	for _, elem_a := range a {
		counter[elem_a] = counter[elem_a] + 1
	}

	// 4. Достаточно пробежаться по элементам массива b и если counter[elem_b] != 0 ЗНАЧИТ И В (b) есть такое занчение
	for _, elem_b := range b {
		if counter[elem_b] == 1 {
			result = append(result, elem_b)
		}
	}
	return result
}

// запусти это в main
//func main() {
//	a := []int{23, 3, 1, 2}
//	b := []int{6, 2, 4, 23}
//	c := week1_28_04_25_04_05_2025.Intersertion(a, b)
//	fmt.Println(c)
//}
