package task3

// 3. Напишите функцию,
// которая вернет кол-во повторений числа в слайсе
// пример:
// input: [1, 2, 1, 2, 2, 5, 5], 2
// output: 3
func countNumberInSlice(numbers []int, search int) int {
	count := 0
	for _, number := range numbers {
		if number == search {
			count++
		}
	}
	return count
	// todo:
	return 0
}
