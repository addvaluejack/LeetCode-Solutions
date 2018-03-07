/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(current *TreeNode, rights []*TreeNode) {
    if current.Right != nil {
        rights = append(rights, current.Right)
    }
    current.Right = current.Left
    current.Left = nil
    if current.Right != nil {
        foo(current.Right, rights)
    } else {
        if len(rights) != 0 {
            current.Right = rights[len(rights)-1]
            rights = rights[:len(rights)-1]
            foo(current.Right, rights)
        }
    }
}

func flatten(root *TreeNode)  {
    if root == nil {
        return
    }
    rights := make([]*TreeNode, 0)
    foo(root, rights)
}
