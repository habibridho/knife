package knife

func Filter(slice []int, filter int) []int {
	result := []int{}
	for i := range slice {
		if slice[i] != filter {
			result = append(result, slice[i])
		}
	}

	return result
}
