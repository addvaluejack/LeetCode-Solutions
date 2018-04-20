/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(current *TreeNode, max_ptr *int) {
    acc := current.Val
    if current.Left != nil {
        foo(current.Left, max_ptr)
        if current.Left.Val > 0 {
            acc += current.Left.Val
        }
    }
    if current.Right != nil {
        foo(current.Right, max_ptr)
        if current.Right.Val > 0 {
            acc += current.Right.Val
        }
    }
    if acc > *max_ptr {
        *max_ptr = acc
    }
    if current.Left != nil && current.Right == nil {
        if current.Left.Val > 0 {
            current.Val += current.Left.Val
        }
    }
    if current.Left == nil && current.Right != nil {
        if current.Right.Val > 0 {
            current.Val += current.Right.Val
        }
    }
    if current.Left != nil && current.Right != nil {
        if current.Left.Val > current.Right.Val {
            if current.Left.Val > 0 {
                current.Val += current.Left.Val
            }
        } else {
            if current.Right.Val > 0 {
                current.Val += current.Right.Val
            }
        }
    }
}

func maxPathSum(root *TreeNode) int {
    max := root.Val
    foo(root, &max)
    return max
}
