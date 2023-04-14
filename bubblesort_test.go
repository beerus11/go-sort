package main

import (
	"fmt"
	"testing"
)

func Test_BubbleSort(t *testing.T) {
	items := []int{3, 7, 6, 3, 1, 9}
	sortedItems := BubbleSort(items)
	fmt.Println(sortedItems)
}
