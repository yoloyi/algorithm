package sorts

import (
	"testing"
)

func TestInsertion(t *testing.T) {
	arr := []int{3, 2, 3, 2, 3, 4, 6, 7, 8, 9345, 2, 34, 1, 4, 5, 6, 7}
	InsertionSort(arr)
	len := len(arr)
	for i := 0; i < len-1; i++ {
		if arr[i] > arr[i+1] {
			t.Error("排序错误")
			break
		}
	}
}
