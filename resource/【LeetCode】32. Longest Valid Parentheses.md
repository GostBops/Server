Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

**Example 1:**
>**Input:** "(()"
**Output:** 2
**Explanation:** The longest valid parentheses substring is "()"

**Example 2:**
>**Input:** ")()())"
**Output:** 4
**Explanation:** The longest valid parentheses substring is "()()"

## 0、题目分析
题目要求给出最长的有效括号序列的长度，因此可以用栈或者动态规划的方法来实现

## 1、栈
我们首先在给定字符串的第一个字符之前插入一个‘）’
之后遍历每一个字符
1. 如果遇到‘（’，就把它的索引压入栈中
2. 如果遇到‘）’，就把栈顶弹出
	1. 弹出后如果栈不为空，就用当前索引减去栈顶储存的索引，得到的值便是该字符串其中一个有效括号序列的长度
	2. 弹出后如果栈为空，就把当前索引压入栈中

```c
class Solution {
public:
    int longestValidParentheses(string s) {
        stack<int> mystack;
        mystack.push(-1);
        int res = 0;
        for(int i = 0; i < s.size(); i++) {
            if(s[i] == '(') {
                mystack.push(i);
            } else {
                mystack.pop();
                if(!mystack.empty()) {
                    res = max(res, i - mystack.top());
                } else {
                    mystack.push(i);
                }
            }
        }
        return res;
    }
};
```

## 2、动态规划
我们维护一个数组DP用来储存对应索引的最大有效括号序列长度
比如dp[5]=4表示给定字符串s从索引0直到索引5的最大有效括号序列长度为4

我们仍然在给定字符串的第一个字符之前插入一个‘）’
然后从第二个字符开始遍历
1. 如果遇到‘（’，我们什么都不做
2. 如果遇到‘）’，设当前索引为i，我们首先考虑前一个索引 j ，根据我们维护的数组DP，如果 DP[ j ]为2，就说明s[ j ]之前包括s[ j ]匹配了两个括号，因此我们需要将s[ i ]和s[j - DP[ j ]]匹配，
	1. 如果s[j - DP[ j ]]是‘（’，就说明匹配成功，将DP[ i ]更新为DP[ j ] + 2，也就是最大有效括号序列长度多了2如果s[j - DP[ j ]]是‘（’，就说明匹配成功，将DP[ i ]更新为DP[ j ] + 2，还要加上DP[j - DP[ j ] - 1]
	2. 如果s[i - DP[ j ]]为‘）’，则匹配失败，那么i之前的序列的匹配工作就结束了，当我们遍历下一个字符i + 1的时候，就是新的一轮括号匹配，这样就把原字符串分隔成一个个子问题来求解

所以一开始我们插入‘）’，意义在于把原字符串作为最初的子问题来求解

```c
class Solution {
public:
    int longestValidParentheses(string s) 
    {
        int res = 0;
        s = ')' + s;
        vector<int> dp(s.length(), 0);
        for(int i = 1; i < s.length(); i++) {
            if(s[i] == ')') {
                if(s[i - 1 - dp[i - 1]] == '(') {
                   dp[i] = dp[i - 1] + 2; 
                } 
                dp[i] += dp[i - dp[i]];
            }
            res = max(res, dp[i]);
        }
        return res;
    }
};
```
