package RadixSort

import (
	"math"
	"strconv"
)

func RadixSort(arr []int)[]int {
	len := len(arr)
	bucket := make([][]int, 10)
	for i := 0; i < 10; i++ {
		bucket[i] = make([]int,0,10)
	}
	maxLen := getMaxNumberLength(arr)
	for i := 0; i < maxLen; i++ {
		for _, v := range arr {
			insertIndex := v / int(math.Pow(10.0, float64(i))) % 10
			bucket[int(insertIndex)] = append(bucket[int(insertIndex)], v)
		}
		arr = make([]int,0,len)
		for index, _ := range bucket {
			arr = append(arr, bucket[index]...)
			bucket[index] = bucket[index][:0]
		}
	}
	return arr
}

func getMaxNumberLength(arr []int) int {
	var max int = 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return len(strconv.Itoa(max))
}
