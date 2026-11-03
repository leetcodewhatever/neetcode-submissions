/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil && root == nil {
		return true
	}
	if root == nil || subRoot == nil || root.Val != subRoot.Val {
		return false
	}

	return isSameTree(root.Left, subRoot.Left) && isSameTree(root.Right, subRoot.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {

  if subRoot == nil {
	return true
  } 

  if root == nil {
	 return false
  }

  if isSameTree(root, subRoot) {
	  return true
  }

  return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
