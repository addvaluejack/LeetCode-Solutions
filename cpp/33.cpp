class Solution {
public:
    int search(vector<int>& nums, int target) {
        if (nums.size() == 0) {
            return -1;
        }
        
        int bottom = 0, top = nums.size();
        
        while (bottom < top) {
            int mid = (bottom+top)/2;
            
            int num = (nums[mid] > nums[0]) == (target > nums[0])
                    ? nums[mid]
                    : target > nums[0] ? INT_MAX : INT_MIN;
            
            if (target < num) {
                top = mid;
            } else if (target > num) {
                bottom = mid + 1;
            } else {
                return mid;
            }
        }
        
        if (target == nums[0]) {
            return 0;
        } else {
            return -1;
        }
    }
};
