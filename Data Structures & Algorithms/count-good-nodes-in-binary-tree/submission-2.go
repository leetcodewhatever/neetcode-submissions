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
  seen := map[int]bool{}
  queue := []*TreeNode{root}
  var first bool = true
  for len(queue) > 0 {
	node := queue[0]
	queue = queue[1:]
   

	if node.Val >= root.Val && !seen[node.Val] {
		n++
        seen[node.Val] = true
		
		if first { // act like you didn't see the root :)
			seen[node.Val] = false
			first = false
		}
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
