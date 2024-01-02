package search

func IndexOf(arr []int, value int) int {
	if isOrdered(arr) {
		return BinarySearch(arr, value)
	}

	return LinearSearch(arr, value)
}

func isOrdered(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}
