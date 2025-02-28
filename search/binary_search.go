package search

import "fmt"

func getInput() {
	var count int
	fmt.Scanln(&count)

	var array []int
	for i := 0; i < count; i++ {
		var number int
		fmt.Scan(&number)
		array = append(array, number)
	}
	var target int
	fmt.Scanln(&target)
	binarySearch(count, array, target)
}

func binarySearch(count int, array []int, target int) {
	left, right := 0, count-1
	middle := (left + right) / 2

	for left <= right {
		middle = (left + right) / 2

		if array[middle] == target {
			fmt.Println(middle)
			return
		}
		if array[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	fmt.Println(left)
}
