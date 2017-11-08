/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import("strconv")

func addPaths(current *TreeNode, path []int, result_ptr *[]string) {
    if current.Left == nil && current.Right == nil {
        var path_str string
        for i := 0; i < len(path)-1; i++ {
            path_str = path_str+strconv.Itoa(path[i])
            path_str = path_str+"->"
        }
        path_str = path_str+strconv.Itoa(path[len(path)-1])
        *result_ptr = append(*result_ptr, path_str)
    }
    if current.Left != nil {
        new_path := make([]int, len(path))
        for i := 0; i < len(path); i++ {
            new_path[i] = path[i]
        }
        new_path = append(new_path, current.Left.Val)
        addPaths(current.Left, new_path, result_ptr)
    }
    if current.Right != nil {
        new_path := make([]int, len(path))
        for i := 0; i < len(path); i++ {
            new_path[i] = path[i]
        }
        new_path = append(new_path, current.Right.Val)
        addPaths(current.Right, new_path, result_ptr)
    }
}

func binaryTreePaths(root *TreeNode) []string {
    result := make([]string, 0)
    if root == nil {
        return result
    }
    path := make([]int, 0)
    path = append(path, root.Val)
    addPaths(root, path, &result)
    
    return result
}
