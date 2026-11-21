/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	mid := indexOf(root.Val, inorder)
    root.Left = buildTree(preorder[1:mid+1], inorder[:mid]) // all the nodes to the left of mid
    root.Right = buildTree(preorder[mid+1:], inorder[mid+1:]) // all the nodes to the right of mid 
    return root
}

func indexOf(val int, list []int) int {
	for i := range list {
		if list[i] == val {
			return i
		}
	}

	return -1
}
