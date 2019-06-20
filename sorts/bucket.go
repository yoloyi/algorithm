package sorts

import (
	"math"
)

// BucketSort 桶排序
func BucketSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	max := arr[1]
	for _, v := range arr {
		max = int(math.Max(float64(v), float64(max)))
	}
	var tmp = make([]int, max+1)
	for _, v := range arr {
		tmp[v]++
	}
	var result []int
	for k, v := range tmp {
		for i := 0; i < v; i++ {
			result = append(result, k)
		}
	}
	arr = result
	copy(arr, result)
}
