package SelectSort

import "DataStructureAlgorithm/BubbleSort"
func SelectSort(arr []int) {
	len := len(arr)
	var index int = 0
	for i := 0; i < len; i++ {
		index = i
		for j := i; j < len; j++ {
			if arr[index] > arr[j] {
				index = j
			}
		}
		BubbleSort.Swap(arr,index,i)
	}
}



