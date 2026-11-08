/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
     stack := []*TreeNode{root}    
	 node := root
	for len(stack) > 0 {
		
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
        // if you shadow the node here, you will lost all the right nodes
		// since the loop destory the node variable.
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// we can also just decrease the counter here
		k--
		if k == 0 {
			return node.Val
		}
        
		// visit the right of every poped left.
        node = node.Right

		}
		return -1 // k is larger then the tree size
}