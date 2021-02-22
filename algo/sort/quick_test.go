package sort_test

import (
	"algo/sort"
	"fmt"
	"testing"
	"time"
)

func TestQuickSort(t *testing.T) {

	defer time.Sleep(time.Second * 2)

	arr := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	fmt.Printf("before quick sort: %v\n", arr)
	sort.QuickSort(arr, 0, len(arr)-1)
	fmt.Printf("after  quick sort: %v\n", arr)
	t.Logf("after  quick sort: %v", arr)
}
