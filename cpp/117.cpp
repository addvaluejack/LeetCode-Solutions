/**
 * Definition for binary tree with next pointer.
 * struct TreeLinkNode {
 *  int val;
 *  TreeLinkNode *left, *right, *next;
 *  TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
 * };
 */
class Solution {
public:
    void connect(TreeLinkNode *root) {
        TreeLinkNode *head;
        TreeLinkNode *prev;
        TreeLinkNode *cur;
        cur = root;
        
        while (cur != NULL) {
            while (cur != NULL) {
                if (cur->left != NULL) {
                    if (prev != NULL) {
                        prev->next = cur->left;
                    } else {
                        head = cur->left;
                    }
                    prev = cur->left;
                }
                if (cur->right != NULL) {
                    if (prev != NULL) {
                        prev->next = cur->right;
                    } else {
                        head = cur->right;
                    }
                    prev = cur->right;
                }
                cur = cur->next;
            }
            
            cur = head;
            prev = NULL;
            head = NULL;
        }
    }
};
