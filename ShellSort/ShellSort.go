package ShellSort

func ShellSort1(nums []int) {
	//外层步长控制
	for step := len(nums) / 2; step > 0; step /= 2 {
		//开始插入排序
		for i := step; i < len(nums); i += step {
			//满足条件则插入
			for j := i - step; j >= 0 && nums[j+step] < nums[j]; j -= step {
				nums[j], nums[j+step] = nums[j+step], nums[j]
			}
		}
	}
}

func ShellSort2(arr []int) {
	len := len(arr)
	for increment := len / 2; increment > 0; increment /= 2 {
		shellSort0(arr,increment)
	}
}

func shellSort0(arr []int, increment int) {
	len := len(arr)
	for step := 0; step < increment; step++ {
		for i := step; i < len; i += increment {
			for j := i - increment; j >= 0 && arr[j+increment] < arr[j]; j -= increment {
				arr[j+increment], arr[j] = arr[j], arr[j+increment]
			}
		}
	}
}
