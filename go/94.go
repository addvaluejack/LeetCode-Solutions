/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(current *TreeNode, result_ptr *[]int) {
    if current == nil {
        return
    }
    if current.Left != nil {
        foo(current.Left, result_ptr)
    }
    *result_ptr = append(*result_ptr, current.Val)
    if current.Right != nil {
        foo(current.Right, result_ptr)
    }
}

func inorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    
    foo(root, &result)
    
    return result
}
