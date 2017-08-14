/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(current *TreeNode, acc int, result_ptr *int) {
    if current == nil {
        if acc > *result_ptr {
            *result_ptr = acc
        }
    } else {
        foo(current.Left, acc+1, result_ptr)
        foo(current.Right, acc+1, result_ptr)
    }
}

func maxDepth(root *TreeNode) int {
    result := 0
    
    foo(root, 0, &result)
    
    return result
}
