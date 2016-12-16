/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
    result := ListNode{0, nil}
    current := &result
    
    for {
        var index int
        flag := true
        for i := 0; i < len(lists); i++ {
            if lists[i] == nil {
                 continue
            }
            if flag == true {
                index = i
                flag = false
            } else if lists[i].Val < lists[index].Val {
                index = i
            }
        }
        if flag == true {
            break;
        } else {
            current.Next = lists[index]
            current = current.Next
            lists[index] = lists[index].Next
        }
    }
    
    return result.Next
}
