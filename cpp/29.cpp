class Solution {
public:
    int divide(int dividend, int divisor) {
        bool dividend_signed = dividend<0?true:false;
        bool divisor_signed = divisor<0?true:false;
        long long result;
        
        if (divisor == 0 || (dividend == INT_MIN && divisor == -1)) {
            return INT_MAX;
        }
        
        if (divisor == INT_MIN && dividend == INT_MIN) {
            return 1;
        }
        
        if (labs(dividend) < labs(divisor)) {
            return 0;
        }
        
        long long dividendx = (long long)dividend;
        long long divisorx = (long long)divisor;
        
        if (dividend_signed && divisor_signed) {
            return func(-dividendx, -divisorx, -divisorx, 1);
        } else if (dividend_signed && !divisor_signed) {
            return -func(-dividendx, divisorx, divisorx, 1);
        } else if (!dividend_signed && divisor_signed) {
            return -func(dividendx, -divisorx, -divisorx, 1);
        } else {
            return func(dividendx, divisorx, divisorx, 1);
        }
    }
    
    long long func(long long dividend, long long divisor, long long acc, long long x) {
        if (dividend > (acc + acc)) {
            return func(dividend, divisor, (acc + acc), x + x);
        } else if (dividend > acc) {
            return x + func(dividend - acc, divisor, divisor, 1);
        } else if (dividend == acc) {
            return 1;
        } else {
            return 0;
        }
    }
};
