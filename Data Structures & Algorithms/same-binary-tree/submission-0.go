/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "slices"
func isSameTree(p *TreeNode, q *TreeNode) bool {
    pList := []int{}
	qList := []int{}

	var traverse func(node *TreeNode, list *[]int)
	traverse = func(node *TreeNode, list *[]int) {

		if node == nil {
			*list = append(*list, math.MinInt)
			return
		}
		*list = append(*list, node.Val)
		traverse(node.Right, list)
		traverse(node.Left, list)
	}
	traverse(p, &pList)
	traverse(q, &qList)

	return slices.Equal(pList, qList)
 
}
