/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo0(current *TreeNode, current_sum int, sum int, result_ptr *int){
    if current_sum+current.Val == sum {
        *result_ptr++
    }
    if current.Left != nil {
        foo1(current.Left, current_sum+current.Val, sum, result_ptr)
        foo0(current.Left, 0, sum, result_ptr)
    }
    if current.Right != nil {
        foo1(current.Right, current_sum+current.Val, sum, result_ptr)
        foo0(current.Right, 0, sum, result_ptr)
    }
}

func foo1(current *TreeNode, current_sum int, sum int, result_ptr *int){
    if current_sum+current.Val == sum {
        *result_ptr++
    }
    if current.Left != nil {
        foo1(current.Left, current_sum+current.Val, sum, result_ptr)
    }
    if current.Right != nil {
        foo1(current.Right, current_sum+current.Val, sum, result_ptr)
    }
}

func pathSum(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }
    result := 0
    foo0(root, 0, sum, &result)
    return result
}
