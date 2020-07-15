package RadixSort

import (
	"fmt"
	"strconv"
	"time"
)

func RadixSort(arr []int)[]int {
	len := len(arr)
	bucket := make([][]int, 10)
	bucketLength := make([]int,10)
	var level1 int = 0
	var level2 int = 0
	for i := 0; i < 10; i++ {
		bucket[i] = make([]int,len,len)
	}
	maxLen := getMaxNumberLength(arr)
	m := 1
	for i := 0; i < maxLen; i++ {
		begin1 := time.Now()
		for _, v := range arr {
			insertIndex := v / m % 10
			i2 := insertIndex
			bucket[i2][bucketLength[i2]] = v
			bucketLength[i2]++
		}
		end1 := time.Now()
		level1+=int(end1.Sub(begin1).Milliseconds())

		begin2 := time.Now()
		arr = make([]int,0,len)
		for index, _ := range bucket {
			arr = append(arr, bucket[index][:bucketLength[index]]...)
			bucketLength[index] = 0
		}
		m *= 10
		end2 := time.Now()
		level2+=int(end2.Sub(begin2).Milliseconds())
	}

	fmt.Println("第一阶段总耗时:",level1,"ms")
	fmt.Println("第二阶段总耗时:",level2,"ms")
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
