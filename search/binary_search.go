package search

func BinarySearch(arr []int, value int) int {
	lo, hi := 0, len(arr)-1
	return binarySearch(arr, value, lo, hi)
}

func binarySearch(arr []int, value, lo, hi int) int {

	if lo == hi {
		if arr[lo] == value {
			return lo
		}

		return -1
	}

	mid := lo + (hi-lo)/2

	if value < arr[mid] {
		return binarySearch(arr, value, lo, mid)
	}
	if value > arr[mid] {
		return binarySearch(arr, value, mid+1, hi)
	}

	return mid

}
