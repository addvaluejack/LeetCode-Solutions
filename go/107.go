/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    var result [][]int
    var queue []*TreeNode
    
    if root == nil {
        return result
    }
    
    queue = append(queue, root)
    for ;len(queue) != 0; {
        var tmp_queue []*TreeNode
        var level_result []int
        
        for i := 0; i < len(queue); i++ {
            level_result = append(level_result, queue[i].Val)
            if queue[i].Left != nil {
                tmp_queue = append(tmp_queue, queue[i].Left)
            }
            if queue[i].Right != nil {
                tmp_queue = append(tmp_queue, queue[i].Right)
            }
        }
        
        result = append(result, level_result)
        queue = tmp_queue
    }
    
    for i := 0; i < len(result)/2; i++ {
        result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
    }
    
    return result
}
