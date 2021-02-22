package sort

import "fmt"

func QuickSort(arr []int, low, high int) {
	if low < high {
		pivot := split1(arr, low, high)
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

// 快排的两种方法。
func split1(arr []int, low, high int) int {
	x := arr[low]
	i, j := low+1, high

	fmt.Printf("low=%d, high=%d\n", low, high)

	for i < j {
		for i < j {
			if x < arr[i] {
				break
			}
			i++
		}
		fmt.Printf("i = %d\n", i)
		for i < j {
			if x < arr[j] {
				break
			}
			j--
		}
		fmt.Printf("j = %d\n", j)

		fmt.Printf("before i=%d,j=%d, a[i]=%d,a[j]=%d\n", i, j, arr[i], arr[j])
		arr[i], arr[j] = arr[j], arr[i]
		fmt.Printf("after  i=%d,j=%d, a[i]=%d,a[j]=%d\n", i, j, arr[i], arr[j])
		break
	}
	arr[i], arr[low] = arr[low], arr[i]
	fmt.Printf("i = %d, arr=%v\n\n", i, arr)
	return i
}
