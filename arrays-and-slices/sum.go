package arraysandslices

func Sum(numbers []int) (result int) {
	for _, value := range numbers {
		result += value
	}
	return result
}
