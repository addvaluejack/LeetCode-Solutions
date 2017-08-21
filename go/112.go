/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(current *TreeNode, sum int, acc int) bool {
    if current.Left == nil && current.Right == nil {
        if acc+current.Val == sum {
            return true
        } else {
            return false
        }
    } else {
        left_bool := false
        right_bool := false
        if current.Left != nil {
            left_bool = foo(current.Left, sum, acc+current.Val)
        }
        if current.Right != nil {
            right_bool = foo(current.Right, sum, acc+current.Val)
        }
        if left_bool == true || right_bool == true {
            return true
        } else {
            return false
        }
    }
}

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    return foo(root, sum, 0)
}
