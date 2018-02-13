/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(t1 *TreeNode, t2 *TreeNode, result *TreeNode) {
    val := 0
    var t1_left *TreeNode
    var t2_left *TreeNode
    var t1_right *TreeNode
    var t2_right *TreeNode
    if t1 != nil {
        val += t1.Val
    } 
    if t2 != nil {
        val += t2.Val
    }
    result.Val = val
    if t1 != nil {
        t1_left = t1.Left
        t1_right = t1.Right
    }
    if t2 != nil {
        t2_left = t2.Left
        t2_right = t2.Right
    }
    if !(t1_left == nil && t2_left == nil) {
        result.Left = new(TreeNode)
        foo(t1_left, t2_left, result.Left)
    }
    if !(t1_right == nil && t2_right == nil) {
        result.Right = new(TreeNode)
        foo(t1_right, t2_right, result.Right)
    }
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    if t1 == nil && t2 == nil {
        return nil
    }
    result := new(TreeNode)
    
    foo(t1, t2, result)
    
    return result
}
