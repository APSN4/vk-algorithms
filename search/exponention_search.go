package search

import "fmt"

func GetInput() {
	var count int
	fmt.Scanln(&count)
	var numbers []int
	for i := 0; i < count; i++ {
		var number int
		fmt.Scan(&number)
		numbers = append(numbers, number)
	}
	var target int
	fmt.Scanln(&target)
	fmt.Println(exponentialSearch(count, numbers, target))
}

func exponentialSearch(count int, numbers []int, target int) (firstPoint int, secondPoint int) {
	border := 1
	for border < count && numbers[border] < target {
		if border > count {
			return border / 2, count
		} else if numbers[border] == target {
			return border, border * 2
		}
		border *= 2
	}
	return border / 2, border
}
