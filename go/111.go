/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(current *TreeNode, result_ptr *int, depth int) {
    if (current.Left == nil && current.Right == nil) && (*result_ptr == -1 || *result_ptr > depth) {
        *result_ptr = depth
    }
    if current.Left != nil {
        foo(current.Left, result_ptr, depth+1)
    }
    if current.Right != nil {
        foo(current.Right, result_ptr, depth+1)
    }
}

func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    result := -1
    foo(root, &result, 1)
    return result
}
