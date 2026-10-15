/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// nil -> 5  7 
	head := &ListNode{}
	prev := head

	for l1 != nil {
	
		sum := l1.Val + l2.Val

		if sum > 9 {
           
		   v1 := &ListNode{
				Val: sum / 10,
		   }

		   v2 := &ListNode{
				Val: sum % 10,
		   }
           
		   v2.Next = v1
           
		   prev.Next = v2
		   prev = prev.Next.Next
		   
		   l1 = l1.Next
		   l2 = l2.Next
		   continue

		}
        

		curr := &ListNode{
			Val: sum,
			Next: nil,
		} 
		
		prev.Next = curr // now, head know the next address
		prev = curr
		
		l1 = l1.Next
		l2 = l2.Next
	}
		
    return head.Next
}
