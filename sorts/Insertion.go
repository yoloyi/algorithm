package sorts

//InsertionSort 直接插入排序
func InsertionSort(arr []int) {
	len := len(arr)
	if len < 2 {
		return
	}
	for i := 1; i < len; i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				break
			}
		}
	}
	return
}

// BinaryInsertionSort 二分插入排序
func BinaryInsertionSort(arr []int) {
	len := len(arr)
	if len < 2 {
		return
	}

	for i := 1; i < len; i++ {
		low := 0
		high := i
		for low <= high {
			middle := (low + high) / 2
			if arr[middle] > arr[i] {
				high = middle - 1
			} else {
				low = middle + 1
			}
		}

		for j := i - 1; j >= low; j-- {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
	return
}
