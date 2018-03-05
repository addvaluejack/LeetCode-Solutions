/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
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
func foo(current *ListNode, length int) *TreeNode {
    result := new(TreeNode)
    i := 0
    new_current := current
    for ; i < length/2; i++ {
        new_current = new_current.Next
    }
    result.Val = new_current.Val
    left_length := i
    if left_length != 0 {
        result.Left = foo(current, left_length)
    }
    right_length := length-i-1
    if right_length != 0 {
        result.Right = foo(new_current.Next, right_length)
    }
    
    return result
}

func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    length := 0
    current := head
    for ; current != nil; current = current.Next {
        length++
    }
    return foo(head, length)
}
