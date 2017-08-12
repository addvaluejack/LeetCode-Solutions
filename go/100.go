/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if (p == nil && q != nil) || (p != nil && q == nil) {
        return false
    }
    if p.Val != q.Val {
        return false
    }
    if (p.Left != nil && q.Left == nil) || (p.Left == nil && q.Left != nil) {
        return false
    }
    if (p.Right != nil && q.Right == nil) || (p.Right == nil && q.Right != nil) {
        return false
    }
    if (p.Left == nil && p.Right == nil) && (q.Left == nil && q.Right == nil) {
        return true
    }
    if isSameTree(p.Left, q.Left) == false {
        return false
    }
    if isSameTree(p.Right, q.Right) == false {
        return false
    }
    return true
}
