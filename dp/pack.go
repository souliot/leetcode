package dp

func PackRec(w, v []int, i int, c int) (res int) {
	if i < 0 || c <= 0 {
		return 0
	}

	// 不放i物品
	A := PackRec(w, v, i-1, c)
	// 放i物品
	B := 0
	if c-w[i] >= 0 {
		B = PackRec(w, v, i-1, c-w[i]) + v[i]
	}
	return Max(A, B)
}

func PackDP(w, v []int, c int) (max int) {
	dp := make([][]int, len(w))
	for i := range dp {
		dp[i] = make([]int, c+1)
	}

	for i := 0; i <= c; i++ {
		if i < w[0] {
			continue
		}
		dp[0][i] = v[0]
	}
	for i := 1; i < len(w); i++ {
		for j := 0; j <= c; j++ {
			if j < w[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = Max(dp[i-1][j-w[i]]+v[i], dp[i-1][j])
			}
		}
	}
	max = dp[len(w)-1][c]
	return
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
