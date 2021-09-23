package link

func PureLinkNode(node *LinkNode) {
	cur := node
	var prev *LinkNode
	cache := make(map[int]struct{})
	for cur != nil {
		if _, ok := cache[cur.Value]; ok {
			cur, prev.Next = cur.Next, cur.Next
			continue
		}
		cache[cur.Value] = struct{}{}
		cur, prev = cur.Next, cur
	}
}
