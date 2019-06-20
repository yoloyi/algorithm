package sorts

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	len := len(arr)
	if len < 2 {
		return
	}
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}
