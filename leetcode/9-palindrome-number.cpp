class Solution {
public:
    bool isPalindrome(int x) {
        string s = to_string(x);
        int l = s.size();
        if (l == 1) {
            return true;
        }
        for(int i = 0; i < l; i++) {
            if (s[i] != s[l-1-i]) {
                return false;
            }
        }

        return true;
    }
};