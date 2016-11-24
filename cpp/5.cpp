class Solution {
public:
    string longestPalindrome(string s) {
        string result = "";
        int longest_string_length = 0;
        
        for(int i = 0; i < s.length();){
            int j = i, k = i;
            
            while(k < s.length() - 1 && s[k] == s[k+1]) {
                k++;
            }
            
            i = k + 1;
            
            while(j > 0 && k < s.length() - 1 && s[j-1] == s[k+1]) {
                j--;
                k++;
            }
            
            
            if (k - j + 1 > result.length()) {
                result = s.substr(j, k - j + 1);
                longest_string_length = k - j + 1;
            }
        }
        
        return result;
    }
};
