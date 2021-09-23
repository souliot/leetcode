package dp

func MaxMul(arr []int) (res []int, max int) {
	cur := arr[0]
	max = cur
	res = []int{cur}
	var gt, gt_max bool
	for i := 1; i < len(arr); i++ {
		cur, gt = MaxAndGT(cur*arr[i], arr[i])
		max, gt_max = MaxAndGT(cur, max)
		if gt {
			if gt_max {
				res = append(res, arr[i])
			}
		} else {
			if gt_max {
				res = []int{arr[i]}
			}
		}
	}

	return
}

func MaxAndGT(a, b int) (max int, gt bool) {
	if a >= b {
		return a, true
	}
	return b, false
}
