class Solution {
    
public:
    vector<int> findSubstring(string s, vector<string>& words) {
        unordered_map<string, int> dict;
        vector<int> result;
        int s_len = s.length();
        int w_len = words[0].length();
        
        for (int i = 0; i < words.size(); i++) {
            dict[words[i]]++;
        }
        
        for (int i = 0; i < w_len; i++) {
            unordered_map<string, int> tdict;
            int count = 0;
            int left = i;
            
            for (int j = i; j <= s_len - w_len; j += w_len) {
                string tword = s.substr(j, w_len);
                
                tdict[tword]++;
                if (dict.count(tword) > 0) {
                    if (tdict[tword] <= dict[tword]) {
                        count++;
                    } else {
                        while (tdict[tword] > dict[tword]) {
                            string leftest_word = s.substr(left, w_len);
                            
                            tdict[leftest_word]--;
                            if (tdict[leftest_word] < dict[leftest_word]) {
                                count--;
                            }
                            left += w_len;
                        }
                    }
                    
                    if (count == words.size()) {
                        result.push_back(left);
                        tdict[s.substr(left, w_len)]--;
                        count--;
                        left += w_len;
                    }
                } else {
                    tdict.clear();
                    count = 0;
                    left = j + w_len;
                }
            }
        }
        
        return result;
    }
};
