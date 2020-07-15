package OddEvenSort

func OddEvenSort(arr []int) {
	var sorted bool = false
	len := len(arr)
	for !sorted {
		sorted = true
		for i:=0;i<len - 1;i+=2 {
			if arr[i] > arr[i+1] {
				arr[i],arr[i+1] = arr[i+1],arr[i]
				sorted = false
			}
		}
		for i := 1; i < len-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i],arr[i+1] = arr[i+1],arr[i]
				sorted = false
			}
		}
	}
}
