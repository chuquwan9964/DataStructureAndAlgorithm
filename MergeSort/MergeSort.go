package MergeSort


func MergeSort(arr []int) {
	mergeSort0(arr,0,len(arr) - 1)
}

func mergeSort0(arr []int,begin,end int)  {
	if end > begin {
		midIndex := begin + (end - begin) / 2
		mergeSort0(arr,begin,midIndex)
		mergeSort0(arr,midIndex+1,end)
		//merge two sorted lists
		mergeTwoSortedList(arr,begin,midIndex,midIndex+1,end)
	}
}

func mergeTwoSortedList(arr []int, a1Begin, a1End, a2Begin, a2End int) {
	
}
