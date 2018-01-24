/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(current *TreeNode, result_ptr *[]int) {
    *result_ptr = append(*result_ptr, current.Val)
    if current.Left != nil {
        foo(current.Left, result_ptr)
    }
    if current.Right != nil {
        foo(current.Right, result_ptr)
    }
}

func preorderTraversal(root *TreeNode) []int {
    result := make([]int, 0)
    
    if root == nil {
        return result
    }
    
    foo(root, &result)
    
    return result
}
