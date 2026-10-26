/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
   
   if root == nil {
		return 0
   }

   queue := []*TreeNode{}
   var count int 

   queue = append(queue, root)

   for len(queue) > 0 {
	   level := len(queue)
	   count++
       
	   var i int
	   for ; i < level; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
	   }
		
   }
   return count

}
