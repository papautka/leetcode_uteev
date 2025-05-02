package __avito_mock_ozon

import "fmt"

func Slice1() {
	a1 := make([]int, 0, 10)                 // len=0 cap=10
	a1 = append(a1, []int{1, 2, 3, 4, 5}...) // len=5 cap=10
	a2 := append(a1, 6)                      // new_struct len=6 cap=10
	a3 := append(a1, 7)                      // new_struct len=6 cap=10
	fmt.Println(a1, a2, a3)
	fmt.Println(len(a1), len(a2), len(a3))
	fmt.Println(cap(a1), cap(a2), cap(a3))
	// [1, 2, 3, 4, 5] [1, 2, 3, 4, 5, 7] [1, 2, 3, 4, 5, 7]
}

func Slice2() {
	a1 := make([]int, 0, 10)                 // len=0 cap=10
	a1 = append(a1, []int{1, 2, 3, 4, 5}...) // len=5 cap=10
	a2 := append(a1, 6)                      // [1, 2, 3, 4, 5, 6]  new_struct len=6 cap=10
	a3 := append(a2, 7)                      // [1, 2, 3, 4, 5, 6, 7]  new_struct len=7 cap=10
	fmt.Println(a1, a2, a3)
	fmt.Println(len(a1), len(a2), len(a3))
	fmt.Println(cap(a1), cap(a2), cap(a3))
	// [1, 2, 3, 4, 5] [1, 2, 3, 4, 5, 6] [1, 2, 3, 4, 5, 6, 7]
}

func Slice3() {
	a1 := make([]int, 0)                     // len=0 cap=0
	a1 = append(a1, []int{1, 2, 3, 4, 5}...) // len=5 cap=5
	a2 := append(a1, 6)                      // [1, 2, 3, 4, 5, 6]  new_struct len=5 cap=6
	a3 := append(a1, 7)                      // [1, 2, 3, 4, 5, 7]  new_struct len=7 cap=6
	fmt.Println(a1, a2, a3)
	fmt.Println(len(a1), len(a2), len(a3))
	fmt.Println(cap(a1), cap(a2), cap(a3))
	// [1, 2, 3, 4, 5] [1, 2, 3, 4, 5, 7] [1, 2, 3, 4, 5, 7]
}
