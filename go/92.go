/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    if m == n {
        return head
    }
    
    var before_m, m_node, n_node *ListNode
    before_m = new(ListNode)
    before_m.Next = head
    
    index := 0
    current := before_m
    
    for ; index+1 < m; {
        current = current.Next
        index++
    }
    
    before_m = current
    m_node = current.Next
    
    next := current.Next
    
    for ; index < n; {
        tmp := next.Next
        next.Next = current
        current = next
        next = tmp
        n_node = current
        index++
    }
    
    before_m.Next = n_node
    m_node.Next = next
    
    if m == 1 {
        return before_m.Next
    } else {
        return head
    }
}
