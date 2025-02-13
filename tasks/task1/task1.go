package task1

// 1. Напишите функцию,
// которая создаст слайс с числами от 1 до n,
// где n - входной параметр (целое число)
// пример:
// input: 5
// output: [1, 2, 3, 4, 5]
func makeSliceOfNumbers(n int) []int {
	s := make([]int, 0, n)

	for i := 1; i <= n; i++ {
		s = append(s, i)
	}
	return s
	// todo:
	return []int{0}
}
