package InsertSort

import "DataStructureAlgorithm/BubbleSort"

func InsertSort(arr []int) {
	len := len(arr)
	var temp int
	for i := 0; i < len; i++ {
		temp = i
		for j:=i-1;j>=0;j-- {
			if arr[temp] < arr[j] {
				BubbleSort.Swap(arr,temp,j)
				temp = j
			}
		}
	}
}
