// Forward declaration of isBadVersion API.
bool isBadVersion(int version);

class Solution {
    int foo(int start, int end) {
        if (start == end) {
            return start;
        }
        if (isBadVersion(start+(end-start)/2)) {
            return foo(start, start+(end-start)/2);
        } else {
            return foo(start+(end-start)/2+1, end);
        }
    }
public:
    int firstBadVersion(int n) {
        int start = 1;
        int end = n;
        return foo(start, end);
    }
};
