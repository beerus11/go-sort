package main

import (
	"fmt"
	"testing"
)

func Test_QuickSort(t *testing.T) {
	items := []int{3, 7, 6, 3, 1, 9}
	Quicksort(items, 0, len(items)-1)
	fmt.Println(items)
}
