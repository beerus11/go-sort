package main

func BubbleSort(items []int) []int {
	if len(items) < 1 {
		return items
	}
	for i := 0; i < len(items); i++ {
		for j := 1; j < len(items)-i; j++ {
			if items[j] < items[j-1] {
				items[j], items[j-1] = items[j-1], items[j]
			}
		}
	}
	return items
}
