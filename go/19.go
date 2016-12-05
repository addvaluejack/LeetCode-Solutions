/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    nodes := make([]*ListNode, 0)
    current := head
    
    for current != nil {
        nodes = append(nodes, current)
        current = current.Next
    }
    
    if n == len(nodes) {
        return head.Next
    } else {
        nodes[len(nodes) - n - 1].Next = nodes[len(nodes) - n].Next
        return head
    }
}
