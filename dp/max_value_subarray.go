package dp

// 1 -2 3 -4 5
func maxProduct(count int, numbers []int, countSubsequence int) int {
	if countSubsequence == 0 {
		return 0
	}

	dpMax := make([][]int, count)
	for i := 0; i < count; i++ {
		dpMax[i] = make([]int, countSubsequence+1)
		for j := 0; j <= countSubsequence; j++ {
			dpMax[i][j] = -9223372036854775808
		}
	}
	dpMin := make([][]int, count)
	for i := 0; i < count; i++ {
		dpMin[i] = make([]int, countSubsequence+1)
		for j := 0; j <= countSubsequence; j++ {
			dpMin[i][j] = 9223372036854775807
		}
	}

	for i := 0; i < count; i++ {
		dpMax[i][0] = 1
		dpMin[i][0] = 1
	}

	for i := 0; i < count; i++ {
		if i < 1 {
			dpMax[i][1] = numbers[i]
			dpMin[i][1] = numbers[i]
		} else {
			if numbers[i] < dpMax[i-1][1] {
				dpMax[i][1] = dpMax[i-1][1]
			} else {
				dpMax[i][1] = numbers[i]
			}
			if numbers[i] < dpMin[i-1][1] {
				dpMin[i][1] = numbers[i]
			} else {
				dpMin[i][1] = dpMin[i-1][1]
			}
		}
	}

	for k := 2; k < countSubsequence+1; k++ {
		for i := 1; i < count; i++ {
			if k > i+1 {
				continue
			}
			dpMax[i][k] = dpMax[i-1][k]
			dpMin[i][k] = dpMin[i-1][k]

			newMax := 0
			newMin := 0
			if numbers[i] >= 0 {
				newMax = numbers[i] * dpMax[i-1][k-1]
				newMin = numbers[i] * dpMin[i-1][k-1]
			} else {
				newMax = numbers[i] * dpMin[i-1][k-1]
				newMin = numbers[i] * dpMax[i-1][k-1]
			}

			if dpMax[i][k] < newMax {
				dpMax[i][k] = newMax
			}
			if dpMin[i][k] > newMin {
				dpMin[i][k] = newMin
			}
		}
	}

	return dpMax[count-1][countSubsequence]
}
