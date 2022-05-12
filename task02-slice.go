package homework

func reverse(input []int64) (result []int64) {
	result = make([]int64, len(input))
	copy(result, input)
	for i := 0; i < len(result)/2; i++ {
		x := len(result) - 1 - i
		result[i], result[x] = result[x], result[i]
	}
	return result
}
