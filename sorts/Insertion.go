package sorts

//InsertionSort 插入排序
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

func main() {

}
