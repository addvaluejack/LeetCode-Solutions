/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
bool findPath(TreeNode* current, TreeNode *target, vector<TreeNode*> path, vector<TreeNode*> &result) {
    if (current == target) {
        path.push_back(current);
        result = path;
        return true;
    } else {
        if (current->left == NULL && current->right == NULL) {
            return false;
        }
        path.push_back(current);
        if (current->left != NULL) {
            if (findPath(current->left, target, path, result)) {
                return true;
            }
        }
        if (current->right != NULL) {
            if (findPath(current->right, target, path, result)) {
                return true;
            }
        }
        return false;
    }
}

class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        vector<TreeNode*> p_path;
        vector<TreeNode*> q_path;
        
        findPath(root, p, p_path, p_path);
        findPath(root, q, q_path, q_path);
        
        for (int i = p_path.size()-1; i >= 0; i--) {
            for (int j = q_path.size()-1; j >= 0 ; j--) {
                if (p_path[i] == q_path[j]) {
                    return p_path[i];
                }
            }
        }
        
        return NULL;
    }
};
