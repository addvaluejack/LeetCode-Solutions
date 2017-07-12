class Solution {
public:
    int firstMissingPositive(vector<int>& nums) {
        for (int i = 0; i < nums.size(); i++) {
            while (nums[i] > 0 && nums[i] <= nums.size() && nums[nums[i]-1] != nums[i]) {
                iter_swap(nums.begin()+i, nums.begin()+nums[i]-1);
            }
        }
        
        int i = 0;
        for (; i < nums.size(); i++) {
            if (nums[i] != i+1) {
                break;
            }
        }
        
        return i+1;
    }
};
