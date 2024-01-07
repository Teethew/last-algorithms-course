package search

func LinearSearch(arr []int, value int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == value {
			return i
		}
	}

	return -1
}
