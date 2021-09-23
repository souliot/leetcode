package link

import "fmt"

type LinkNode struct {
	Value int
	Next  *LinkNode
}

func (m *LinkNode) String() (res string) {
	res = fmt.Sprintf("%d", m.Value)
	for m.Next != nil {
		m = m.Next
		res = fmt.Sprintf("%s->%d", res, m.Value)
	}
	return
}

func (m *LinkNode) Append(v int) {
	for m.Next != nil {
		m = m.Next
	}
	m.Next = &LinkNode{
		Value: v,
	}
}

func (m *LinkNode) Unshift(v int) {
	old := *m
	*m = LinkNode{
		Value: v,
	}
	(*m).Next = &old
}
