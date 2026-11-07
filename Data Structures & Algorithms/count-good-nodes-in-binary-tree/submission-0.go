/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
   if root == nil {
		return 0
   } 

  var n int 
  queue := []*TreeNode{root}
  for len(queue) > 0 {
	node := queue[0]
	queue = queue[1:]
    
	if node.Val >= root.Val {
		n++
	}

	if node.Right != nil {
		queue = append(queue, node.Right)
	}

	if node.Left != nil {
		queue = append(queue, node.Left)
	}

  }

  return n
}
