package link

import "testing"

func TestReverseAppend(t *testing.T) {
	n := &LinkNode{Value: 5}
	n.Append(6)
	n.Append(7)
	n.Append(8)
	t.Log(n)

	r := Reverse(n)
	t.Log(r)
}
func TestReverse(t *testing.T) {
	n := &LinkNode{Value: 5}
	n.Unshift(6)
	n.Unshift(7)
	n.Unshift(8)
	t.Log(n)

	r := Reverse(n)
	t.Log(r)
}
