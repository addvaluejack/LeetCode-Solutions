class Solution {
public:
    int trap(vector<int>& A) {
        int left = 0; int right = A.size()-1;
        int result = 0;
        int max_left = 0, max_right = 0;
        while(left <= right){
            if(A[left] <= A[right]){
                if(A[left] >= max_left) {
                    max_left = A[left];
                } else {
                    result += max_left-A[left];
                }
                left++;
            }
            else{
                if(A[right] >= max_right) {
                    max_right = A[right];
                } else {
                    result += max_right-A[right];
                }
                right--;
            }
        }
        return result;
    }
};
