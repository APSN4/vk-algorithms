package hash_tables

// adeuuuuio -> 4 (uuuu)
func maxCountSymbol(line string) int {
	symbols := make(map[string]int)
	for i := 0; i < len(line); i++ {
		letter := string(line[i])
		if _, ok := symbols[letter]; ok {
			symbols[letter] += 1
		} else {
			symbols[letter] = 1
		}
	}
	maxValue := 0
	for _, v := range symbols {
		if v > maxValue {
			maxValue = v
		}
	}
	return maxValue
}
