class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
        vector<int> result(digits);
        int carry = 0, i = digits.size()-1;
        
        carry = (digits[i]+1)/10;
        result[i] = (digits[i]+1)%10;
        i--;
        
        while(i >= 0 && carry != 0) {
            carry = (digits[i]+1)/10;
            result[i] = (digits[i]+1)%10;
            i--;
        }
        
        if(i == -1 && carry != 0) {
            result.insert(result.begin(), 1);
        }
        
        return result;
    }
};
