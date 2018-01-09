/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(current *TreeNode, sum int, path []int, result_ptr *[][]int) {
    current_sum := 0
    path = append(path, current.Val)
    for i := 0; i < len(path); i++ {
        current_sum += path[i]
    }
    if current_sum == sum && current.Left == nil && current.Right == nil {
        tmp_path := make([]int, len(path))
        for i := 0; i < len(path); i++ {
            tmp_path[i] = path[i]
        }
        *result_ptr = append(*result_ptr, tmp_path)
    }
    if current.Left != nil {
        foo(current.Left, sum, path, result_ptr)
    }
    if current.Right != nil {
        foo(current.Right, sum, path, result_ptr)
    }
}

func pathSum(root *TreeNode, sum int) [][]int {
    result := make([][]int, 0)
    path := make([]int, 0)
    if root != nil {
        foo(root, sum, path, &result)        
    }
    return result
}
