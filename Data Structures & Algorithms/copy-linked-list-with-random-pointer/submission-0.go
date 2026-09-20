/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	newHead := &Node{}
	headStart := newHead
	oldNewMap := make(map[*Node]*Node)

    if head == nil {
        return head
    }
    if head.Next == nil {
        node := &Node{}
        if head.Random != nil {
            node.Val = head.Val
            node.Next = nil
            node.Random = node
        }else{
             node.Val = head.Val
        }
        return node
    }
	for head.Next != nil {
		//
		// newNextTwo := &Node{}
		newNextOne := &Node{}
		// nextNode := head.Next

		// handle next
		_, exists := oldNewMap[head.Next]
		if !exists {
			newNextOne.Val = head.Next.Val
			newHead.Next = newNextOne
			oldNewMap[head] = newHead
			oldNewMap[head.Next] = newNextOne
		} else {
			newHead.Next = oldNewMap[head.Next]
			oldNewMap[head] = newHead
			oldNewMap[head.Next] = oldNewMap[head.Next]
		}

		// use mapping to extract the random pointer

		// oldNewMap[head] = newHead
		// oldNewMap[head.Next] = newNextOne

		//

		_, exists = oldNewMap[head.Random]
		if !exists {
			// if it doesn't exist, create the node

			if head.Random != nil {
				random := &Node{}
				random.Val = head.Random.Val
				newHead.Random = random
				oldNewMap[head.Random] = random
			} else {
				newHead.Random = nil
				oldNewMap[head.Random] = nil
			}

		} else {
			random := oldNewMap[head.Random]
			newHead.Random = random
            oldNewMap[head.Random] = random
		}

		newHead.Val = head.Val

		head = head.Next
		newHead = newHead.Next
	}

	if head.Next == nil && head.Random != nil {
		newHead.Random = oldNewMap[head.Random]
	}
	return headStart
}
