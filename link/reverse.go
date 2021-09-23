package link

func Reverse(head *LinkNode) (r *LinkNode) {
	cur := head
	var prev *LinkNode
	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}
	return prev
}
