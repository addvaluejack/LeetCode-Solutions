class Solution {
public:
    int reverse(int x) {
        bool neg = false;
        int result = 0;
        
        if(x < 0) {
            neg = true;
            x = -x;
        }
        
        while(x != 0){
            int t1 = result*10;
            if(t1/10 != result) {
                result = 0;
                break;
            }
            int t2 = t1 + x%10;
            if(t2 - x%10 != t1) {
                result = 0;
                break;
            }
            result = t2;
            x = x/10;
        }
        
        if(neg) {
            return -result;
        } else {
            return result;
        }
    }
};
