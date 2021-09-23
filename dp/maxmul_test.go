package dp

import "testing"

func TestMul(t *testing.T) {
	arr := []int{3, 1, 7, -2, 4, -6, 8}

	res, max := MaxMul(arr)
	t.Log(res, max)
}
