/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func reverseIntSlice(input []int) []int {
    for i := 0; i < len(input)/2; i++ {
        input[i], input[len(input)-1-i] = input[len(input)-1-i], input[i]
    }
    return input
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    current_layer := make([]*TreeNode, 0)
    current_layer_index := 0
    
    if root == nil {
        return result
    }
    current_layer = append(current_layer, root)
    for ; len(current_layer) != 0; current_layer_index++ {
        tmp := make([]int, 0)
        next_layer := make([]*TreeNode, 0)
        for i := 0; i < len(current_layer); i++ {
            tmp = append(tmp, current_layer[i].Val)
            if current_layer[i].Left != nil {
                next_layer = append(next_layer, current_layer[i].Left)
            }
            if current_layer[i].Right != nil {
                next_layer = append(next_layer, current_layer[i].Right)
            }
        }
        if current_layer_index%2 == 1 {
            result = append(result, reverseIntSlice(tmp))
        } else {
            result = append(result, tmp)
        }
        current_layer = next_layer
    }
    
    return result
}
