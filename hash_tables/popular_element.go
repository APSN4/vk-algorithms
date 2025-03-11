package hash_tables

func PopularElement(count int, elements []int) int {
	frequency := make(map[int]int)

	for _, k := range elements {
		if _, exists := frequency[k]; exists {
			frequency[k] += 1
		} else {
			frequency[k] = 1
		}

		if val, exists := frequency[k]; exists && val > count/2 {
			return k
		}
	}

	return -1
}
