class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        vector<int> result(2, -1);
        vector<int> no_found(2, -1);
        
        if (nums.size() == 0 || target < nums[0] || target > nums[nums.size()-1]) {
            return no_found;
        }
        
        int bottom = 0, top = nums.size();
        // find starting position
        result[0] = bottom;
        if (target > nums[bottom]) {
            while (bottom < top) {
                int mid = (bottom+top)/2;
                
                if (target > nums[mid]) {
                    bottom = mid + 1;
                } else if (target < nums[mid]) {
                    top = mid;
                } else {
                    if (target > nums[mid-1]) {
                        result[0] = mid;
                        break;
                    } else {
                        top = mid;
                    }
                }
            }
            if (nums[result[0]] != target) {
                return no_found;
            }
        }
        
        bottom = 0, top = nums.size();
        // find ending position
        result[1] = top-1;
        if (target < nums[top-1]) {
            while (bottom < top) {
                int mid = (bottom+top)/2;
                
                if (target > nums[mid]) {
                    bottom = mid + 1;
                } else if (target < nums[mid]) {
                    top = mid;
                } else {
                    if (target < nums[mid+1]) {
                        result[1] = mid;
                        break;
                    } else {
                        bottom = mid + 1;
                    }
                }
            }
            if (nums[result[0]] != target) {
                return no_found;
            }
        }
        
        return result;
    }
};
