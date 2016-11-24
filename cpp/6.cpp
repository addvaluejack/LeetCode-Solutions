class Solution {
public:
    string convert(string s, int numRows) {
        vector<string> Matrix;
        
        if(numRows == 1){
            return s;
        }
        
        for(int i = 0; i < numRows; i++){
            Matrix.push_back("");
        }
        
        for(int i = 0; i < s.length(); i++) {
            int rem = i % (numRows*2 - 2);

            if(rem < numRows) {
                Matrix[rem] += s[i];
            } else {
                for(int j = 0; j < numRows; j++) {
                    if (j == (numRows*2 - 2) - rem){
                        Matrix[j] += s[i];
                    }
                }
            }
        }
        
        string result = "";
        for(int i = 0; i < numRows; i++){
            result.append(Matrix[i]);
        }
        
        return result;
    }
};
