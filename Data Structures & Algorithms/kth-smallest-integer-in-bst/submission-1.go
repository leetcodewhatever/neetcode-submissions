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
     h := []int{}
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
		h = append(h, node.Val)
        node = node.Right

		}
		return h[k-1]
}
