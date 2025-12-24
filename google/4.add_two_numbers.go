package google

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) addToHead(node *ListNode) *ListNode {
	if head == nil {
		return node
	}
	if node != nil {
		(head).Next = node
		head = (head).Next
		return head
	}
	return nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//r1 := reverseList(l1)
	//r2 := reverseList(l2)

	ans := addNodes(l1, l2)
	//return reverseList(ans)
	return ans
}

func addNodes(a *ListNode, b *ListNode) *ListNode {
	carry := false
	var head, curr *ListNode
	for a != nil || b != nil {
		var aVal, bVal int
		if a != nil {
			aVal = a.Val
			a = a.Next
		}
		if b != nil {
			bVal = b.Val
			b = b.Next
		}
		total := aVal + bVal
		if carry {
			total++
		}
		sum := total % 10
		carry = total >= 10

		node := &ListNode{Val: sum}

		if head == nil {
			head = node
			curr = node
		} else {
			curr.Next = node
			curr = curr.Next
		}
	}

	if carry {
		curr.Next = &ListNode{Val: 1}
	}

	return head

}

func reverseList(head *ListNode) *ListNode {
	var prev, curr, next *ListNode

	curr = head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
