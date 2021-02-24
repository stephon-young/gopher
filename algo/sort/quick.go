package sort

func QuickSort(arr []int, low, high int) {
	if low < high {
		pivot := split2(arr, low, high)
		QuickSort(arr, low, pivot-1)
		QuickSort(arr, pivot+1, high)
	}
}

// 快排的两种方法。
func split1(arr []int, low, high int) int {
	x := arr[low]
	i, j := low, high

	//fmt.Printf("low=%d, high=%d\n", low, high)

	for i < j {
		for i < j {
			if x > arr[j] {
				break
			}
			j--
		}
		if i != j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		//fmt.Printf("left i=%d,j=%d, arr=%v\n", i, j, arr)

		for i < j {
			if x < arr[i] {
				break
			}
			i++
		}
		if i != j {
			arr[i], arr[j] = arr[j], arr[i]
			j--
		}
		//fmt.Printf("right i=%d,j=%d, arr=%v\n", i, j, arr)
	}
	arr[i] = x
	//fmt.Printf("i = %d, arr=%v\n\n", i, arr)
	return i
}

func split2(arr []int, low, high int) int {
	k := low
	for i := low + 1; i <= high; i++ {
		if arr[i] < arr[low] {
			k++
			if k != i {
				arr[k], arr[i] = arr[i], arr[k]
			}
		}
	}
	arr[low], arr[k] = arr[k], arr[low]
	return k
}
