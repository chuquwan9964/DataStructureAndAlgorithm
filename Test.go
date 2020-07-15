package main

import (
	"DataStructureAlgorithm/RadixSort"
	"fmt"
	"math/rand"
	"time"
)
func main() {
	var baseNum = 10000000
	slice := make([]int, 0)
	//var slice = []int{9,8,7,6,5,4,3,2,1,10,19,18,17,16,15,14,13,12,11}
	for i := baseNum; i > 0; i-- {
		slice = append(slice, rand.Intn(baseNum * 100))
		//slice = append(slice, i)
	}
	begin := time.Now()
	RadixSort.RadixSort(slice)
	end := time.Now()
	fmt.Println(end.Sub(begin).Milliseconds(),"ms")
}

func main1() {
	fmt.Println(len("hello world"))
}