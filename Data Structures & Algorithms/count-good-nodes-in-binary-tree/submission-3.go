/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
   var dfs func(node *TreeNode, maxSoFar int) int
   dfs = func(node *TreeNode, maxSoFar int) int{

		if node == nil {
			return 0
		}

		count := 0
		if node.Val >= maxSoFar {
			count = 1
            maxSoFar = node.Val
		}

		return count + dfs(node.Right, maxSoFar) + dfs(node.Left, maxSoFar)
   }

  return dfs(root, root.Val)
}
