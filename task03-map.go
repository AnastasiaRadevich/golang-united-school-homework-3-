package homework

func sortMapValues(input map[int]string) (result []string) {
	result = make([]string, len(input))
	for i := 0; i < len(input); i++ {
		result[i] = input[i]
	}
	return result
}
