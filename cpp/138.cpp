/**
 * Definition for singly-linked list with a random pointer.
 * struct RandomListNode {
 *     int label;
 *     RandomListNode *next, *random;
 *     RandomListNode(int x) : label(x), next(NULL), random(NULL) {}
 * };
 */
class Solution {
public:
    RandomListNode *copyRandomList(RandomListNode *head) {
        if (head == NULL) {
            return NULL;
        }
        
        vector<RandomListNode *> old_list;
        vector<RandomListNode *> new_list;
        
        RandomListNode *new_head = new RandomListNode(0);
        new_head->label = head->label;
        
        RandomListNode *old_current = head;
        RandomListNode *new_current = new_head;
        old_list.push_back(old_current);
        new_list.push_back(new_current);
        
        while (old_current->next != NULL) {
            new_current->next = new RandomListNode(0);
            new_current->next->label = old_current->next->label;
            new_current = new_current->next;
            old_current = old_current->next;
            old_list.push_back(old_current);
            new_list.push_back(new_current);
        }
        
        old_current = head;
        new_current = new_head;
        
        while (old_current != NULL) {
            if (old_current->random != NULL) {
                int i = 0;
                for (; i < old_list.size(); ++i) {
                    if (old_list[i] == old_current->random) {
                        break;
                    }
                }
                new_current->random = new_list[i];
            }
            old_current = old_current->next;
            new_current = new_current->next;
        }
        
        return new_head;
    }
};
