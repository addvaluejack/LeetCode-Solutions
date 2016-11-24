class Solution {
public:
    vector<int> findDisappearedNumbers(vector<int>& nums) {
        vector<int> result(nums.size(), 0);
        for(int i = 0; i < nums.size(); i++) {
            result[nums[i] - 1] = -1;
        }
        for(int i = 0; i < nums.size(); i++) {
            if(result[i] == 0) {
                result[i] = i + 1;
            }
        }
        int i = 0;
        while(true) {
            if(i == result.size()) {
                break;
            }
            if(result[i] == -1) {
                result.erase(result.begin()+i);
            } else {
                i++;
            }
        }
        
        return result;
    }
};
