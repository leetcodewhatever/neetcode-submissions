/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    visited := make(map[*ListNode]struct{});
 
    for head != nil{
        _, exists := visited[head]
        if exists{
            return true
        }else{
            visited[head] = struct{}{}
            head = head.Next
        }
    }

    return false
    
}