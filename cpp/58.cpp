class Solution {
public:
    int lengthOfLastWord(string s) {
        int result = 0;
        int pre_length = 0;
        
        for (int i = 0; i < s.length(); i++) {
            if (s[i] == ' ') {
                if (result != 0) {
                    pre_length = result;
                }
                result = 0;
                continue;
            }
            result++;
        }
        
        if (result == 0) {
            return pre_length;
        }
        return result;
    }
};
