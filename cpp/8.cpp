class Solution {
public:
    int myAtoi(string str) {
        int result = 0;
        bool neg = false;
        bool exceeded = false;
        
        int i = 0;
        
        while(str[i] == ' ') {
            i++;
        }
        
        if(str[i] == '-') {
            neg = true;
            i++;
        } else if(str[i] == '+') {
            i++;
        }
        
        for (; i < str.length(); i++) {
            if (str[i] < '0' || str[i] > '9') {
                break;
            } else {
                int t1 = result*10 + (str[i] - '0');
                if((t1 - (str[i] - '0')) / 10 != result || t1 < result) {
                    exceeded = true;
                    break;
                } else {
                    result = t1;
                }
            }
        }
        
        if(exceeded) {
            if(neg) {
                return INT_MIN;
            } else {
                return INT_MAX;
            }
        } else {
            if(neg) {
                return -result;
            } else {
                return result;
            }
        }
    }
};
