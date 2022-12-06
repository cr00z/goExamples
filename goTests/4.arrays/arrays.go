package arrays

func Sum(ops []int) int {
	var result int
	for _, op := range ops {
		result += op
	}
	return result
}

func SumAll(arrs ...[]int) []int {
	var result []int
	for _, arr := range arrs {
		result = append(result, Sum(arr))
	}
	return result
}

func SumAll2(arrs ...[]int) []int {
	result := make([]int, len(arrs))
	for i, arr := range arrs {
		result[i] = Sum(arr)
	}
	return result
}
