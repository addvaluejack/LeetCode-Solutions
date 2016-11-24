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
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        int inc = 0, sum;
        ListNode* fake_head = new ListNode(-1);
        ListNode* current = fake_head;
        
        while(true) {
            if(l1 == NULL && l2 == NULL) {
                if(inc != 0){
                    current->next = new ListNode(inc);
                }
                break;
            } else if(l1 != NULL && l2 != NULL) {
                sum = l1->val + l2->val + inc;
                inc = 0;
                current->next = new ListNode(sum%10);
                inc = sum/10;
                current = current->next;
                l1 = l1->next;
                l2 = l2->next;
            } else if(l1 != NULL && l2 == NULL) {
                sum = l1->val + inc;
                inc = 0;
                current->next = new ListNode(sum%10);
                inc = sum/10;
                current = current->next;
                l1 = l1->next;
            } else if(l1 == NULL && l2 != NULL) {
                sum = l2->val + inc;
                inc = 0;
                current->next = new ListNode(sum%10);
                inc = sum/10;
                current = current->next;
                l2 = l2->next;
            }
        }
        
        return fake_head->next;
    }
};
