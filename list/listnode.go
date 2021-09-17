package list

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func (m *ListNode) String() (res string) {
	res = fmt.Sprintf("%d", m.Value)
	for m.Next != nil {
		m = m.Next
		res = fmt.Sprintf("%s->%d", res, m.Value)
	}
	return
}

func (m *ListNode) Append(v int) {
	for m.Next != nil {
		m = m.Next
	}
	m.Next = &ListNode{
		Value: v,
	}
}

func (m *ListNode) Unshift(v int) {
	old := *m
	*m = ListNode{
		Value: v,
	}
	(*m).Next = &old
}
