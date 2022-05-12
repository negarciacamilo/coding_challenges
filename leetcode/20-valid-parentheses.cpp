/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

    Open brackets must be closed by the same type of brackets.
    Open brackets must be closed in the correct order.



Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true

Example 3:

Input: s = "(]"
Output: false



Constraints:

    1 <= s.length <= 104
    s consists of parentheses only '()[]{}'.


*/

class Solution
{
public:
    bool isValid(string s)
    {
        if (s.size() <= 1)
        {
            return false;
        }
        map<char, char> closing;
        closing[')'] = '(';
        closing['}'] = '{';
        closing[']'] = '[';
        std::stack<char> symbolStack;
        for (auto c : s)
        {
            if (closing[c])
            {
                if (symbolStack.empty())
                {
                    return false;
                }
                char curr = symbolStack.top();
                if (curr != closing[c])
                {
                    return false;
                }
                else
                {
                    symbolStack.pop();
                }
            }
            else
            {
                symbolStack.push(c);
            }
        }
        if (!symbolStack.empty())
            return false;
        return true;
    }
};