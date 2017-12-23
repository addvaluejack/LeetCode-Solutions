/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return make([]*TreeNode, 0)
    }
    
    return generateSubTrees(1, n)
}

func generateSubTrees(s int, e int) []*TreeNode {
    result := make([]*TreeNode, 0)
    if s > e {
        result = append(result, nil)
        return result
    }
    
    for i := s; i <= e; i++ {
        leftSubTrees := generateSubTrees(s, i-1)
        rightSubTrees := generateSubTrees(i+1, e)
        for _, left := range leftSubTrees {
            for _, right := range rightSubTrees {
                root := new(TreeNode)
                root.Val = i
                root.Left = left
                root.Right = right
                result = append(result, root)
            }
        }
    }
    
    return result
}
