package list

func Reverse(head *ListNode) (r *ListNode) {
	cur := head
	var prev *ListNode
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}
	return prev
}
