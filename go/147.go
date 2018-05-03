/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    pre_head := new(ListNode)
    pre_head.Next = head
    unsorted := head.Next
    head.Next = nil
    for ; unsorted != nil; {
        current := pre_head
        for ; current.Next != nil && current.Next != unsorted && current.Next.Val < unsorted.Val; current = current.Next {}
        tmp := unsorted.Next
        unsorted.Next = current.Next
        current.Next = unsorted
        unsorted = tmp
    }
    return pre_head.Next
}
