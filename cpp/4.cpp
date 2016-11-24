class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        vector<int> combined_nums;
        int size1 = nums1.size(), size2 =nums2.size();
        
        for(int i = 0; i < size1+size2; i++){
            if(nums1.empty()) {
                combined_nums.push_back(nums2.back());
                nums2.pop_back();
                continue;
            } else if (nums2.empty()) {
                combined_nums.push_back(nums1.back());
                nums1.pop_back();
                continue;
            }
            if(nums1.back() > nums2.back()){
                combined_nums.push_back(nums1.back());
                nums1.pop_back();
            } else {
                combined_nums.push_back(nums2.back());
                nums2.pop_back();
            }
        }
        
        if((size1+size2)%2 == 1){
            return float(combined_nums[(size1+size2)/2]);
        } else {
            return float(combined_nums[(size1+size2)/2]*0.5 + combined_nums[(size1+size2)/2 - 1]*0.5);
        }
    }
};
