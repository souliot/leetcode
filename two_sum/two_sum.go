package two_sum

func twoSumMap(arr []int, target int) (rets [][]int) {
	cache := make(map[int]int)
	rets = make([][]int, 0)
	for i, v := range arr {
		if j, ok := cache[target-v]; ok {
			rets = append(rets, []int{i, j})
		} else {
			cache[v] = i
		}
	}
	return
}

func twoSumLoop(arr []int, target int) (rets [][]int) {
	rets = make([][]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				rets = append(rets, []int{i, j})
			}
		}
	}
	return
}

func twoSumSELoop(arr []int, target int) (rets [][]int) {
	rets = make([][]int, 0)
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				rets = append(rets, []int{i, j})
			}
		}
	}
	return
}
