package sorts

import "fmt"

//InsertionSort 简单插入排序
func InsertionSort(arr []int) {
	len := len(arr)
	if len < 2 {
		return
	}
	for i := 1; i < len; i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				j = j - 1
			} else {
				break
			}
		}
		arr[j+1] = key
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

// ShellSort 希尔排序
// 也称递减增量排序算法，是插入排序的一种更高效的改进版本。希尔排序是非稳定排序算法。
func ShellSort(arr []int) {
	len := len(arr)
	step := len / 2

	for step > 0 {
		for i := 0; i < len-1; i += step {
			for j := i; j >= 0; j -= step {
				if arr[j] > arr[j+step] {
					arr[j], arr[j+step] = arr[j+step], arr[j]
				}
			}
		}
		step = step / 2
	}
	fmt.Println(arr)
}
