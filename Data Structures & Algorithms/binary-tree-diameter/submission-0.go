/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
   stack := []*TreeNode{}

   stack = append(stack, root)
   var diameter int
   var count int

   for len(root) > 0 {
	   last := len(stack) - 1
	   node := stack[last]
	   count ++
	   stack = stack[:last]

	   if node.Right != nil {
		  stack = append(stack, node.Right)
	   }

	   if node.Left != nil {
		  stack = append(stack, node.Left)
	   }

      
	  if node.Right == nil && node.Left == nil {
		 
	  }


   }
}
