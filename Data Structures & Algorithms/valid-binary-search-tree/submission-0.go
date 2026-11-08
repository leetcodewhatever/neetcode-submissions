/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type item struct {
	node *TreeNode
	low int
	high int
}
func isValidBST(root *TreeNode) bool {
	stack := []item{item{root, math.MinInt64, math.MaxInt64}}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

        node, low, high := top.node, top.low, top.high 
        if node == nil {

		} 
		if node.Val <= low || node.Val >= high {
			return false
		}
  
        if node.Left != nil {
			// the idea here is that we pass the current node value as the high
			// the upcoming nodes on the left should always be less then it
			stack = append(stack, item{node.Left, low, node.Val})
		}


		if node.Right != nil {
			// the idea here is that we pass the node value as the low
			// then on the right, the nodes should be always bigger then that low.
			stack = append(stack, item{node.Right, node.Val, high})
		}

	}
	return true
}

