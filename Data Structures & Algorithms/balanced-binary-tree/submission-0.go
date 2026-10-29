/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
   
   var height func(node *TreeNode) int
   height = func (node *TreeNode) int {
		if node == nil {
			return 0
		}

		lh := height(node.Left)
		if lh == -1 {
			return -1
		}

		rh := height(node.Right)
		if rh == -1 {
			return -1 
		}
		
		diff := lh - rh
        if diff > 1 || diff < -1 {
			return -1
		}

		return 1 + max(lh, rh)
   }

   h := height(root)

   return h != -1

}
