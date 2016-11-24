class Solution {
public:
    string longestPalindrome(string s) {
        string result = "";
        
        for(int i = 0; i < s.length(); i++) {
            for(int j = 1; j <= s.length() - i; j++) {
                string tmp = s.substr(i, j);
                bool flag = true;
                for(int k = 0; k < static_cast<int>(tmp.length())/2; k++) {
                    if(tmp[k] != tmp[static_cast<int>(tmp.length()) - 1 - k]) {
                        flag = false;
                    }
                }
                if(flag == true && result.length() < tmp.length()){
                    result = tmp;
                }
            }
        }
        
        return result;
    }
};
