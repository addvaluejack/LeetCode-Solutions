/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func foo(nums []int, start int, end int, current *TreeNode) {
    max := nums[start]
    max_index := start
    for i := start; i <= end; i++ {
        if nums[i] > max {
            max = nums[i]
            max_index = i
        }
    }
    current.Val = max
    if max_index > start {
        current.Left = new(TreeNode)
        foo(nums, start, max_index-1, current.Left)
    }
    if max_index < end {
        current.Right = new(TreeNode)
        foo(nums, max_index+1, end, current.Right)
    }
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    } else {
        result := new(TreeNode)
        foo(nums, 0, len(nums)-1, result)
        return result
    }
