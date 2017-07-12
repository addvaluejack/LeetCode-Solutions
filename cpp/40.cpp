class Solution {
public:
    int check_combinations(vector<int>& candidates, int target, vector<int> count) {
        int acc = 0;
        
        for (int i = 0; i < candidates.size(); i++) {
            acc += candidates[i] * count[i];
            if (acc > target) {
                return 1;
            }
        }
        
        if (acc == target) {
            return 0;
        } else if (acc > target) {
            return 1;
        } else {
            return -1;
        }
    }
    
    void foo(vector<int>& candidates, int target, vector<int> count, int count_index, vector<vector<int>>& results) {
        if (count_index == candidates.size()) {
            if (check_combinations(candidates, target, count) == 0) {
                vector<int> result;
                for (int i = 0; i < candidates.size(); i++) {
                    for (int j = 0; j < count[i]; j++) {
                        result.push_back(candidates[i]);
                    }
                }
                results.push_back(result);
            }
            return;
        }
        
        int stride = 1;
        while (count_index + stride < candidates.size() && candidates[count_index] == candidates[count_index+stride]) {
            stride++;
        }
        
        for (int i = 0; i < stride+1; i++) {
            count[count_index] = i;
            if (check_combinations(candidates, target, count) != 1) {
                foo(candidates, target, count, count_index+stride, results);
            }
        }
    }
    
    vector<vector<int>> combinationSum2(vector<int>& candidates, int target) {
        vector<vector<int>> results;
        vector<int> count(candidates.size(), 0);
        
        sort(candidates.begin(), candidates.begin()+candidates.size());
        foo(candidates, target, count, 0, results);
        
        return results;
    }
};
