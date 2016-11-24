class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int max_length = 0;
        vector<int> stop_flag(s.length(), 0);
        for (int i = 0; i < s.length(); i++){
            int current_length = 1;
            int j = i - 1;
            int current_stop_flag = stop_flag[j];
            
            for (; j >= current_stop_flag; j--){
                if (s[j] != s[i]){
                    current_length++;
                } else {
                    break;
                }
                if (stop_flag[j] > current_stop_flag) {
                    current_stop_flag = stop_flag[j];
                }
            }
            stop_flag[i] = j + 1;
            if(current_length > max_length) {
                max_length = current_length;
            }
        }
        
        return max_length;
    }
};
