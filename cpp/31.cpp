class Solution {
public:
    void nextPermutation(vector<int>& nums) {
        for (int i = nums.size() - 2; i > -1; i--) {
            int candidate_index = i + 1;
            int candidate_value = nums[candidate_index];
            
            for (int j = i + 1; j < nums.size(); j++) {
                if (nums[j] > nums[i] && nums[j] <= candidate_value) {
                    candidate_index = j;
                    candidate_value = nums[j];
                }
            }
            
            if (candidate_value > nums[i]) {
                swap(nums[i], nums[candidate_index]);
                sort(nums.begin() + i + 1, nums.end());
                return;
            }
        }
        sort(nums.begin(), nums.end());
    }
};
