package main

import (
	"algo/sort"
	"fmt"
	"time"
)

func main() {
	defer time.Sleep(time.Second * 2)

	mergeSort()
}

func quickSort() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Printf("before quick sort: %v\n", arr)
	sort.QuickSort(arr, 0, len(arr)-1)
	fmt.Printf("after  quick sort: %v\n", arr)
}

func mergeSort() {
	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Printf("before merge sort: %v\n", arr)
	sort.MergeSort(arr, 0, len(arr)-1)
	fmt.Printf("after  merge sort: %v\n", arr)
}
