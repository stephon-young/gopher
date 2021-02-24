package sort

func MergeSort(arr []int, left, right int) {
	if left < right {
		mid := (left + right) / 2
		MergeSort(arr, left, mid)
		MergeSort(arr, mid+1, right)
		Merge(arr, left, mid, right)
	}
}

func Merge(arr []int, left, mid, right int) {
	i, j := left, mid+1

	temp := make([]int, 0, right-left+1)
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}
	for i <= mid {
		temp = append(temp, arr[i])
		i++
	}
	for j <= right {
		temp = append(temp, arr[j])
		j++
	}

	for i := 0; i < right-left+1; i++ {
		arr[left+i] = temp[i]
	}
}
