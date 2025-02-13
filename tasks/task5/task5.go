package task5

import "tasks/tasks/task4"

// 5. Напишите функцию,
// которая вернет кол-во четных чисел в слайсе
// пример:
// input: [1, 2, 1, 2, 2, 5, 5, 4, 6]
// output: 5
func countEvenNumbers(numbers []int) int {
	count := 0
	for _, number := range numbers {
		if task4.IsEven(number) {
			count++
		}
	}
	return count
	// todo:
	//return 0
}
