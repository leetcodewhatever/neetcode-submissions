/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    var diameter int

    var height func(node *TreeNode) int
    height = func(node *TreeNode) {
		if node == nil {
			return 0
		}

		left := height(node.Left) // return 0
		right := height(node.Right) // return 0
        
		if left + right > diameter {
			diameter = left + right
		}
        
		// report back the size of the current node.
		return 1 + max(left, right)

	}
	 
    diam := height(root)
	return diam

}
