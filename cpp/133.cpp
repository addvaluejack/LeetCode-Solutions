/**
 * Definition for undirected graph.
 * struct UndirectedGraphNode {
 *     int label;
 *     vector<UndirectedGraphNode *> neighbors;
 *     UndirectedGraphNode(int x) : label(x) {};
 * };
 */
#include <map>

class Solution {
    map<int, UndirectedGraphNode *> node_map;
    map<int, UndirectedGraphNode *>::iterator it;
    
    void foo(UndirectedGraphNode *cloned, UndirectedGraphNode *original) {
        for(int i = 0; i < original->neighbors.size(); i++) {
            it = node_map.find(original->neighbors[i]->label);
            if(it == node_map.end()) {
                # create new node
                UndirectedGraphNode *new_node = new UndirectedGraphNode(original->neighbors[i]->label);
                node_map.insert(pair<int, UndirectedGraphNode *>(original->neighbors[i]->label, new_node));
                cloned->neighbors.push_back(new_node);
                foo(new_node, original->neighbors[i]);
            } else {
                cloned->neighbors.push_back(it->second);
            }
        }
    }
public:
    UndirectedGraphNode *cloneGraph(UndirectedGraphNode *node) {
        if(node == NULL) {
            return NULL;
        }
        UndirectedGraphNode *result = new UndirectedGraphNode(node->label);
        node_map.insert(pair<int, UndirectedGraphNode *>(node->label, result));
        foo(result, node);
        return result;
    }
};
