/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getDepth(current *TreeNode) int {
    if current.Left == nil && current.Right == nil {
        current.Val = 1
    }
    if current.Left == nil && current.Right != nil {
        right_depth := getDepth(current.Right)
        if right_depth == -1 {
            current.Val = -1
        } else if right_depth > 1 {
            current.Val = -1
        } else {
            current.Val = right_depth+1
        }
    }
    if current.Left != nil && current.Right == nil {
        left_depth := getDepth(current.Left)
        if left_depth == - 1 {
            current.Val = -1
        } else if left_depth > 1 {
            current.Val = -1
        } else {
            current.Val = left_depth+1
        }
    }
    if current.Left != nil && current.Right != nil {
        left_depth := getDepth(current.Left)
        right_depth := getDepth(current.Right)
        if left_depth == -1 || right_depth == -1 {
            current.Val = -1
        } else if left_depth - right_depth > 1 || right_depth - left_depth > 1 {
            current.Val = -1
        } else if left_depth > right_depth {
            current.Val = left_depth+1
        } else {
            current.Val = right_depth+1
        }
    }
    return current.Val
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    getDepth(root)
    if root.Val != -1 {
        return true
    } else {
        return false
    }
}
