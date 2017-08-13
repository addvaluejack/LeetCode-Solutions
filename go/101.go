/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    
    var queue []*TreeNode
    queue = append(queue, root)
    
    for ;; {
        var tmp []*TreeNode
        flag := false
        
        for i := 0; i < len(queue)/2; i++ {
            if queue[i] == nil && queue[len(queue)-1-i] == nil  {
                continue
            }
            if queue[i] == nil && queue[len(queue)-1-i] != nil {
                return false
            }
            if queue[i] != nil && queue[len(queue)-1-i] == nil  {
                return false
            }
            if queue[i].Val != queue[len(queue)-1-i].Val {
                return false
            }
        }
        for i := 0; i < len(queue); i++ {
            if queue[i] != nil {
                tmp = append(tmp, queue[i].Left)
                tmp = append(tmp, queue[i].Right)
                flag = true
            }
        }
        queue = tmp
        if flag == false {
            break
        }
    }
    
    return true
}
