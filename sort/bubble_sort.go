package sort

func BubbleSort(arr []int) {
	last := len(arr)
	bubbleSort(arr, last)
}

func bubbleSort(arr []int, last int) {
	if last <= 1 {
		return
	}

	for i := 0; i < last-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	bubbleSort(arr, last-1)
}
