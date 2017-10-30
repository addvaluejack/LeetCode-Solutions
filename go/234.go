/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    var rev *ListNode
    rev = nil
    fast, slow := head, head
    
    for ; fast != nil && fast.Next != nil; {
        fast = fast.Next.Next
        tmp := rev
        rev = slow
        slow = slow.Next
        rev.Next = tmp
    }
    if fast != nil{
        slow = slow.Next
    }
    for ; rev != nil && rev.Val == slow.Val; {
        slow, rev = slow.Next, rev.Next
    }
    
    return rev == nil
}
