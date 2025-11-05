package task2

// 1. Напишите функцию,
// которая создаст слайс с числами от n до 1,
// где n - входной параметр (целое число)
// пример:
// input: 5
// output: [5, 4, 3, 2, 1]
func makeSliceOfNumbersReverse(n int) []int {
	// todo:
	return []int{0}
}{
		result := []int{}
		for i := n; i >= 1; i--{
			result = append(result, i)
		}
		return result
}
