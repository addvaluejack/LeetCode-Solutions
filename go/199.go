/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    result := make([]int, 0)
    if root == nil {
        return result
    }
    current_layer := make([]*TreeNode, 0)
    current_layer = append(current_layer, root)
    for ; len(current_layer) != 0; {
        result = append(result, current_layer[len(current_layer)-1].Val)
        tmp_layer := make([]*TreeNode, 0)
        for i := 0; i < len(current_layer); i++ {
            if current_layer[i].Left != nil {
                tmp_layer = append(tmp_layer, current_layer[i].Left)
            }
            if current_layer[i].Right != nil {
                tmp_layer = append(tmp_layer, current_layer[i].Right)
            }
        }
        current_layer = tmp_layer
    }
    return result
}
