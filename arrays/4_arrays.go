package arrays

import "fmt"

func GetInput() {
	var count int
	fmt.Scanln(&count)

	var containers []int
	for i := 0; i < count; i++ {
		var codeContainer int
		fmt.Scan(&codeContainer)

		containers = append(containers, codeContainer)
	}

	searchLastEvenContainer(containers)
}

func searchLastEvenContainer(containers []int) {
	storageIndex := -1

	for i, el := range containers {
		if el%2 == 0 || el == 0 {
			storageIndex = i
		}
	}

	fmt.Println(containers[storageIndex])
}
