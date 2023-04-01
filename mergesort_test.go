package main

import (
	"fmt"
	"testing"
)

func Test_MergeSort(t *testing.T) {
	items := []int{3, 7, 6, 3, 1, 9}
	sortedItems := mergeSort(items)
	fmt.Println(sortedItems)
}
