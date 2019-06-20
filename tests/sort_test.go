package tests

import (
	"testing"

	"github.com/yoloyi/algorithm/sorts"
	"github.com/yoloyi/algorithm/utils"
)

var (
	arr  = []int{3, 2, 3, 2, 3, 4, 6, 7, 8, 9345, 2, 34, 1, 4, 5, 6, 7}
	arr2 = make([]int, len(arr))
)

func TestInsertionSort(t *testing.T) {
	copy(arr2, arr)
	sorts.InsertionSort(arr)
	for _, v := range arr {
		if _, err := utils.Contain(v, arr2); err != nil {
			t.Error(err)
			break
		}
	}
	len := len(arr)
	for i := 0; i < len-1; i++ {
		if arr[i] > arr[i+1] {
			t.Error("排序错误")
			break
		}
	}
}

func TestBinaryInsertionSort(t *testing.T) {
	copy(arr2, arr)
	sorts.InsertionSort(arr)
	for _, v := range arr {
		if _, err := utils.Contain(v, arr2); err != nil {
			t.Error(err)
			break
		}
	}
	len := len(arr)
	for i := 0; i < len-1; i++ {
		if arr[i] > arr[i+1] {
			t.Error("排序错误")
			break
		}
	}
}

func TestBubbleSort(t *testing.T) {
	copy(arr2, arr)
	sorts.BubbleSort(arr)
	for _, v := range arr {
		if _, err := utils.Contain(v, arr2); err != nil {
			t.Error(err)
			break
		}
	}
	len := len(arr)
	for i := 0; i < len-1; i++ {
		if arr[i] > arr[i+1] {
			t.Error("排序错误")
			break
		}
	}
}
