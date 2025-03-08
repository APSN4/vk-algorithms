package sort

/*
numbers := []int{2, 7, 0, 1, 2, 4}
numbersPointer := &numbers
sort.QuickSort(numbersPointer, 0, len(numbers))
fmt.Println(numbers)
*/

func QuickSort(unsorted *[]int, start int, end int) {
	if start >= end {
		return
	}

	iPivot := partition(*unsorted, start, end-1)
	QuickSort(unsorted, start, iPivot)
	QuickSort(unsorted, iPivot+1, end)
}

func partition(unsorted []int, start int, end int) int {
	pivot := unsorted[end]
	iPivot := start

	for i := start; i < end; i++ {
		if unsorted[i] <= pivot {
			unsorted[i], unsorted[iPivot] = unsorted[iPivot], unsorted[i]
			iPivot++
		}
	}

	unsorted[iPivot], unsorted[end] = unsorted[end], unsorted[iPivot]

	return iPivot
}
