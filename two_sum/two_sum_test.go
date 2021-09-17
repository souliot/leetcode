package two_sum

import "testing"

func TestTwoSumMap(t *testing.T) {
	arr := []int{1, 2, 4, 8, 3, 6, 0, 9, 11, 18, 15}

	ret := twoSumMap(arr, 14)
	t.Log(ret)
}

func TestTwoSumLoop(t *testing.T) {
	arr := []int{1, 2, 4, 8, 3, 6, 0, 9, 11, 18, 15}

	ret := twoSumLoop(arr, 14)
	t.Log(ret)
}
