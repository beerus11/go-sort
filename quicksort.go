package main

func Quicksort(arr []int, low, high int) {
	if low < high {
		pivot := parition(arr, low, high)
		Quicksort(arr, low, pivot-1)
		Quicksort(arr, pivot+1, high)
	}
}

func parition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i += 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
