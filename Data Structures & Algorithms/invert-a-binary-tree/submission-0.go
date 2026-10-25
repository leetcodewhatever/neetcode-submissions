/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
   // i can go level by level
   // swap left to be right and right to be left
   // since it is level by level, we can use a bfs with a queue
   if root == nil {
		return nil
   }
   queue := []*TreeNode{}

   queue = append(queue, root)
   
   for len(queue) > 0 {
		// peek and then pop the value
		// swap the left and right
		// push the queue
        // fmt.Println(queue)
		node := queue[0]
		fmt.Println(node.Val)
		queue = queue [1:]
		
		node.Left, node.Right = node.Right, node.Left

		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

   }
   return root
}
