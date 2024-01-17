package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, lo, hi int) {
	if lo >= hi || lo < 0 {
		return
	}

	lt, gt := partition(arr, lo, hi)

	quickSort(arr, lo, lt-1)
	quickSort(arr, gt+1, hi)
}

func partition(arr []int, lo, hi int) (lt, gt int) {
	pivot := arr[(lo+hi)/2]

	// Equal, lesser and greater index
	eq := lo
	lt = lo
	gt = hi

	for eq <= gt {

		if arr[eq] < pivot {

			arr[eq], arr[lt] = arr[lt], arr[eq]
			lt += 1
			eq += 1

		} else if arr[eq] > pivot {

			arr[eq], arr[gt] = arr[gt], arr[eq]
			gt -= 1

		} else {

			eq += 1
		}
	}

	return lt, gt
}
