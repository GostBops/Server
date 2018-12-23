Given a non negative integer number **num**. For every numbers **i** in the range **0 ≤ i ≤ num** calculate the number of 1's in their binary representation and return them as an array.

**Example 1:**

>**Input:** 2
**Output:** [0,1,1]

Example 2:

>**Input:** 5
**Output:** [0,1,1,2,1,2]

## 0、题目分析
题目要求给出每个数字的二进制的1的个数，基本思路是使用动态规划和位操作，一开始没有想到好的位操作，硬凑出来了一种解法，结果耗时只超过一半提交。然后在讨论区看到一种解法，非常自然，几乎超过百分百的提交，下面介绍这种解法：

解法遵循两个判断：
1. 当一个数是偶数时，因为最低位为0，所以它的1的位数和它的二分之一的1的位数是一样多的。
`ans[i] = ans[i>>1]`
2. 当一个数是奇数时，肯定比它上一个数多一个1，因为上一个数是偶数，最低位为0，加1不进位。
`ans[i] = ans[i-1]+1`

## 1、代码实现
```c
class Solution {
public:
    vector<int> countBits(int num) {
        vector<int> ans(num+1, 0);
        for(int i = 1; i <= num; ++i) {
            if(i%2 == 0) 
                ans[i] = ans[i>>1];
            else 
                ans[i] = ans[i-1]+1;            
        }
        return ans;
    }
};
```


最开始的辣鸡算法：
```c
class Solution {
public:
    vector<int> countBits(int num) {
        vector<int> bits(num + 1);
        bits[0] = 0;
        for(int i = 1; i <= num; i++) {
            int n = i ^ (i - 1);
            if(n > i - 1) {
                bits[i] = 1;
            } else {
                n = (n + 1) / 2;
                bits[i] = bits[n] + bits[i - n];
            }
        }
        return bits;
    }
};
```
