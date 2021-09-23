package dp

import "testing"

var (
	w = []int{5, 5, 3, 4, 3}
	v = []int{400, 200, 300, 350, 500}
)

func TestPackRec(t *testing.T) {
	cap := 10
	max := PackRec(w, v, 4, cap)
	t.Log(max)
}

func TestPackDP(t *testing.T) {
	cap := 10
	max := PackDP(w, v, cap)
	t.Log(max)
}
