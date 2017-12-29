/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(postorder []int, postindex_ptr *int, subinorder []int) * TreeNode {
    result := new(TreeNode)
    if len(subinorder) == 1 {
        result.Val = subinorder[0]
        (*postindex_ptr)--
    } else {
        rootindex := 0
        for ; rootindex < len(subinorder); rootindex++ {
            if subinorder[rootindex] == postorder[*postindex_ptr] {
                break
            }
        }
        result.Val = postorder[*postindex_ptr]
        (*postindex_ptr)--
        if rootindex < len(subinorder)-1 {
            result.Right = foo(postorder, postindex_ptr, subinorder[rootindex+1:len(subinorder)])
        }
        if rootindex > 0 {
            result.Left = foo(postorder, postindex_ptr, subinorder[0:rootindex])
        }
    }
    
    return result
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(postorder) == 0 {
        return nil
    }
    postindex := len(postorder)-1
    return foo(postorder, &postindex, inorder)
}
