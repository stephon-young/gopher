package main

import (
	"algo/sort"
	"fmt"
	"time"
)

func main() {
	defer time.Sleep(time.Second * 2)

	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Printf("before quick sort: %v\n", arr)
	sort.QuickSort(arr, 0, len(arr)-1)
	fmt.Printf("after  quick sort: %v\n", arr)
}
