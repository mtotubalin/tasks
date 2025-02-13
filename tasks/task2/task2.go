package task2

// 1. Напишите функцию,
// которая создаст слайс с числами от n до 1,
// где n - входной параметр (целое число)
// пример:
// input: 5
// output: [5, 4, 3, 2, 1]
func makeSliceOfNumbersReverse(n int) []int {
	s := make([]int, 0, n)

	for i := n; i >= 1; i-- {
		s = append(s, i)
	}
	return s
	// todo:
	return []int{0}
}
