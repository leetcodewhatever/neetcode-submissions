/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type item struct{
	node *TreeNode
	maxSoFar int
}

func goodNodes(root *TreeNode) int {
  // stack version
  var n int
  stack :=  []item{item{root, root.Val}}
  for len(stack) > 0 {
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]


    node, maxSoFar := top.node, top.maxSoFar
	if node.Val >= maxSoFar {
       n++
	   maxSoFar = node.Val
	}

	if node.Right != nil {
		stack = append(stack, item{node.Right, maxSoFar})
	}

	if node.Left != nil {
		stack = append(stack, item{node.Left, maxSoFar})
	}

  } 
  return n
}

