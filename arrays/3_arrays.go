package arrays

import "fmt"

func getInput() {
	var count int
	fmt.Scanln(&count)

	records := make([]int, count)
	var value int
	for i := 0; i < count; i++ {
		fmt.Scan(&value)
		records[i] = value
	}
	maxRecord(count, records)
}

func maxRecord(count int, records []int) {
	maxValue := 0
	for i := 0; i < count; i++ {
		if records[i] > maxValue {
			maxValue = records[i]
		}
	}
	fmt.Println(maxValue)
}
