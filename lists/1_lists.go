package lists

import (
	"fmt"
)

func GetInput() {
	var count int
	fmt.Scanln(&count)

	var array []int
	for i := 0; i < count; i++ {
		var number int
		fmt.Scan(&number)
		array = append(array, number)
	}
	minAbsolute(count, array)
}

func minAbsolute(count int, array []int) {
	firstNumber := array[len(array)-2]
	secondNumber := array[len(array)-1]
	minSum := firstNumber + secondNumber

	for i := 0; i < count-1; i++ {
		checkCond := array[i+1] - array[i]
		if checkCond < minSum {
			minSum = checkCond
			firstNumber = array[i]
			secondNumber = array[i+1]
		}
	}
	fmt.Println(firstNumber, secondNumber)
}
