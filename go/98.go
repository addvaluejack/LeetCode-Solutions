/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(current *TreeNode, min int, max int) bool {
    if current == nil {
        return true
    }
    if current.Val <= min || current.Val >= max {
        return false
    }
    return foo(current.Left, min, current.Val) && foo(current.Right, current.Val, max)
}

func isValidBST(root *TreeNode) bool {
    return foo(root, -2147483649, 2147483648)
}
