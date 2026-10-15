/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    
	if head.Next == nil && n >= 1 {
		return nil
	}

    curr := head
	var prev *ListNode 

    count := 0
	
	for head != nil {
		
		if count != n {
			prev = curr
			curr = curr.Next
			count++
			continue
		}
        
         
        prev.Next = curr.Next
		break

	} 

	return head
}
