class Solution {
public:
    bool isPalindrome(int x) {
        int len = 0;
        
        if(x == 0) {
            return true;
        }
        
        if(x < 0) {
            return false;
        }
        
        int tmp_x = x;
        while(tmp_x != 0) {
            len++;
            tmp_x = tmp_x/10;
        }
        
        tmp_x = x;
        if(len % 2 == 0) {
            int half = 0;
            
            for(int i = 0; i < len/2; i++) {
                half = half*10 + tmp_x%10;
                tmp_x = tmp_x/10;
            }
            
            if(tmp_x - half == 0) {
                return true;
            } else {
                return false;
            }
        } else {
            int half = 0;
            
            for(int i = 0; i < len/2; i++) {
                half = half*10 + tmp_x%10;
                tmp_x = tmp_x/10;
            }
            
            tmp_x = tmp_x/10;
            if(tmp_x - half == 0) {
                return true;
            } else {
                return false;
            }
        }
    }
};
