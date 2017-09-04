/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
        ListNode *currA = headA;
        ListNode *currB = headB;
        bool loopA = false, loopB = false;
        
        if (headA == NULL || headB == NULL) {
            return NULL;
        }
        
        while (true) {
            if (currA == currB) {
                return currA;
            }
            if (currA == NULL || currB == NULL) {
                break;
            }
            if (loopA == false && currA->next == NULL) {
                currA = headB;
                loopA = true;
            } else {
                currA = currA->next;
            }
            if (loopB == false && currB->next == NULL) {
                currB = headA;
                loopB = true;
            } else {
                currB = currB->next;
            }
        }
        
        return NULL;
    }
};
