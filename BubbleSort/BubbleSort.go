package BubbleSort

func BubbleSort(arr []int)  {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := 0; j < len-i - 1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr,j,j+1)
			}
		}
	}
}

func Swap(arr []int, index1, index2 int) {
	var temp int = arr[index1]
	arr[index1] = arr[index2]
	arr[index2] = temp
}

