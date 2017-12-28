/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(preorder []int, preindex_ptr *int, subinorder []int) * TreeNode {
    result := new(TreeNode)
    if len(subinorder) == 1 {
        result.Val = subinorder[0]
        (*preindex_ptr)++
    } else {
        rootindex := 0
        for ; rootindex < len(subinorder); rootindex++ {
            if subinorder[rootindex] == preorder[*preindex_ptr] {
                break
            }
        }
        result.Val = preorder[*preindex_ptr]
        (*preindex_ptr)++
        if rootindex > 0 {
            result.Left = foo(preorder, preindex_ptr, subinorder[0:rootindex])
        }
        if rootindex < len(subinorder)-1 {
            result.Right = foo(preorder, preindex_ptr, subinorder[rootindex+1:len(subinorder)])
        }
    }
    
    return result
}


func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    preindex := 0
    return foo(preorder, &preindex, inorder)
}
