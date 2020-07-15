package QuickSort

func QuickSort(arr []int) {
	quickSort0(arr,0,len(arr) - 1)
}

func quickSort0(arr []int, begin, end int) {
	var l,r,baseVal int
	l = begin
	r = end
	baseVal = arr[l]
	for l < r {
		for l < r && baseVal <= arr[r] {
			r--
		}
		arr[l] = arr[r]
		for l < r && baseVal >= arr[l] {
			l++
		}
		arr[r] = arr[l]
	}
	arr[r] = baseVal
	if begin < r-1 {
		quickSort0(arr,begin,r - 1)
	}
	if r+1 < end {
		quickSort0(arr,r+1,end)
	}
}

func Quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	Quick2Sort(values[:head])
	Quick2Sort(values[head+1:])
}
