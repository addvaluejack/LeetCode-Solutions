/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func constructTree(node_ptr *TreeNode, nums []int) {
    node_ptr.Val = nums[len(nums)/2]
    nums_left := nums[0:len(nums)/2]
    if len(nums_left) != 0 {
        node_ptr.Left = new(TreeNode)
        constructTree(node_ptr.Left, nums_left)
    }
    nums_right := nums[len(nums)/2+1:len(nums)]
    if len(nums_right) != 0 {
        node_ptr.Right = new(TreeNode)
        constructTree(node_ptr.Right, nums_right)
    }
    
}

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    
    result := new(TreeNode)
    constructTree(result, nums)
    
    return result
}
