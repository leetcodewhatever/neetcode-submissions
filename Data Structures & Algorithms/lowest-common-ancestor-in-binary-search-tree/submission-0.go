/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {

	if root == nil || root == p || root == q{
		return root
	}
	
    right := lowestCommonAncestor(root.Right, p,q)
	left := lowestCommonAncestor(root.Left, p,q)
    

    if right != nil && left != nil{
       return root 
	}

	if right != nil {
		return right
	}

	return left

}
