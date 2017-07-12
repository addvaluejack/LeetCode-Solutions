class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        if (target <= nums[0]) {
            return 0;
        }
        if (target > nums[nums.size()-1]) {
            return nums.size();
        }
        if (target == nums[nums.size()-1]) {
            return nums.size()-1;
        }
        
        int bottom = 0, top = nums.size();
        while (bottom < top) {
            int mid = (bottom+top)/2;
            
            if (target > nums[mid-1] && target <= nums[mid]) {
                return mid;
            }
            if (target > nums[mid]) {
                bottom = mid+1;
            } else {
                top = mid;
            }
        }
        
        return -1;
    }
};
