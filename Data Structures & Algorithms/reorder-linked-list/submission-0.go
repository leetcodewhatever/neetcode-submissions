/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
    
	second := slow.Next
	slow.Next = nil

	var prev *ListNode

	for second != nil {
		nxt := second.Next
		
		second.Next = prev 
		
		prev = second
		second = nxt
	}
    

	first, second := head, prev


	for second != nil && first != nil {
       firstNext, secondNext := first.Next, second.Next
       first.Next = second
	   second.Next = firstNext
	   first = firstNext
	   second = secondNext
	}

	// Now we have two lists
	// head is pointing to the start of the first list
	// prev is pointing the the start of the reserved list



}