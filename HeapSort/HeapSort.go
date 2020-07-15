package HeapSort

import (
	"DataStructureAlgorithm/BubbleSort"
)
//the key is the adjustment function
func HeapSort(arr []int) {
	makeBigTop(arr)
	for i := len(arr) - 1; i > 0; i-- {
		BubbleSort.Swap(arr,i,0)
		adjust(arr,0,i - 1)
	}
}

func adjust(arr []int, rootIndex int,end int) {
	lI := rootIndex*2 + 1
	if lI > end {
		return
	}
	var maxIndex  = lI
	if maxIndex + 1 <= end && arr[maxIndex+1] > arr[maxIndex]{
		maxIndex ++
	}
	if arr[maxIndex] > arr[rootIndex] {
		BubbleSort.Swap(arr,maxIndex,rootIndex)
		adjust(arr,maxIndex,end)
	}
}

func makeBigTop(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjust(arr,i,len(arr) - 1)
	}
}
