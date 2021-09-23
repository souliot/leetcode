package link

import "testing"

func TestPureLink(t *testing.T) {
	n := &LinkNode{
		Value: 1,
	}
	n.Append(1)
	n.Append(2)
	n.Append(3)
	n.Append(3)
	n.Append(3)
	n.Append(4)
	n.Append(5)
	n.Append(5)
	n.Append(7)
	t.Log(n)
	PureLinkNode(n)
	t.Log(n)
}
