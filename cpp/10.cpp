class Solution {
    typedef struct State {
        bool asterisk;
        char element;
        struct State* next;
    } State;
    
    State* createDFA(string p) {
        if(p.length() == 0) {
            return NULL;
        }
        int i = 0;
        State* head = new State;
        head->asterisk = false;
        head->element = p[i];
        head->next = NULL;
        i++;
        if(i < p.length() && p[i] == '*') {
            head->asterisk = true;
            i++;
        }
        State* current = head;
        while(i < p.length()) {
            current->next = new State;
            current->next->asterisk = false;
            current->next->element = p[i];
            current->next->next = NULL;
            i++;
            if(i < p.length() && p[i] == '*') {
                current->next->asterisk = true;
                i++;
            }
            current = current->next;
        }
        return head;
    }
    
    bool walkThroughDFA(string s, int index, State* current) {
        if(current == NULL) {
            if(index == s.length()) {
                return true;
            } else {
                return false;
            }
        } else {
            if(index == s.length()) {
                if(current->asterisk == true) {
                    if(current->next == NULL) {
                        return true;
                    } else {
                        return walkThroughDFA(s, index, current->next);
                    }
                } else {
                    return false;
                }
            } else {
                if(current->asterisk == true) {
                    if(current->next != NULL) {
                        if(current->next->element == s[index] || current->next->element == '.' || current->next->asterisk == true) {
                            if(walkThroughDFA(s, index, current->next) == true) {
                                return true;
                            } 
                        }
                    }
                    if(current->element == s[index] || current->element == '.') {
                        return walkThroughDFA(s, index+1, current);
                    } else {
                        return false;
                    }
                } else {
                    if(current->element == s[index] || current->element == '.') {
                        return walkThroughDFA(s, index+1, current->next);
                    } else {
                        return false;
                    }
                }
            }
        }
    }
    
public:
    bool isMatch(string s, string p) {
        State* head = createDFA(p);
        
        return walkThroughDFA(s, 0, head);
    }
};
