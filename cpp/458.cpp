class Solution {
public:
    int poorPigs(int buckets, int minutesToDie, int minutesToTest) {
        int base  = minutesToTest / minutesToDie;
        
        if (buckets == 1){
            return 0;
        }
        
        int exp = 1;
        while(true) {
            if (pow((base + 1), exp) >= buckets) {
                return exp;
            } else {
                exp++;
            }
        }
    }
};
