class Solution {
public:
    int longestValidParentheses(string s) {
        vector<int> milestones(s.length());

        for (int i = 0; i < s.length(); i++) {
            if (s[i] == '(') {
                milestones[i] = 0;
            }
            else if (s[i] == ')') {
                if (i - 1 >= 0) {
                    if (s[i - 1] == '(') {
                        milestones[i] = 1;
                        if (i - 2 >= 0 && s[i - 2] == ')') {
                            milestones[i] += milestones[i - 2];
                        }
                    }
                    else if (s[i - 1] == ')') {
                        if (milestones[i - 1] == 0) {
                            milestones[i] = 0;
                        }
                        else {
                            if (i - milestones[i-1]*2 - 1 < 0) {
                                milestones[i] = 0;
                            }
                            else {
                                if (s[i - milestones[i-1]*2 - 1] == '(') {
                                    milestones[i] = 1 + milestones[i - 1];
                                    if (i - milestones[i-1]*2 - 2 && s[i - milestones[i-1]*2 - 2] == ')') {
                                        milestones[i] += milestones[i - milestones[i-1]*2 - 2];
                                    }
                                }
                                else {
                                    milestones[i] = 0;
                                }
                            }
                        }
                    }
                }
                else {
                    milestones[i] = 0;
                }
            }
        }

        int result = 0;
        for (int j = 0; j < s.length(); j++) {
            if (result < milestones[j]) {
                result = milestones[j];
            }
        }

        return result * 2;
    }
};
