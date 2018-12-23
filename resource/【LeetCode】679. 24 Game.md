You have 4 cards each containing a number from 1 to 9. You need to judge whether they could operated through *, /, +, -, (, ) to get the value of 24.

**Example 1:**
>**Input:** [4, 1, 8, 7]
**Output:** True
**Explanation:** (8-4) * (7-1) = 24

**Example 2:**
>**Input:** [1, 2, 1, 2]
**Output:** False

**Note:**
- The division operator / represents real division, not integer division. For example, 4 / (1 - 2/3) = 12.
- Every operation done is between two numbers. In particular, we cannot use - as a unary operator. For example, with [1, 1, 1, 1] as input, the expression -1 - 1 - 1 - 1 is not allowed.
- You cannot concatenate numbers together. For example, if the input is [1, 2, 1, 2], we cannot write this as 12 + 12.


## 0、题目分析
这次的题目很有意思，是小时候经常玩的凑24点游戏，也就是说用四个数字经过加减乘除的运算后，结果为24。

那么最简单直接的方式便是暴力地将这四个数字的所有组合遍历一遍，直到发现24的运算结果

## 1、算法设计
那么应该怎么遍历所有组合呢？
我们先来看两个数字的情况，比如1和2，它们的所有组合分别为1+2=3, 1-2=-1，2-1=1， 2x1=2，2/1 = 2, 1/2 = 0.5，我们用（1,2）代替1和2的所有组合
如果再加一个数字，变成1、2和3的组合呢，那么显然1和2的每一个组合（1,2），都可以和3实现一组新的加减乘除的组合，即（3,3）（-1,3）（1,3）（2,3）（2,3）（0.5,3），我们用（（1,2），3）代替（1,2）和3的所有组合
那么1、2和3的所有组合便是（（1,2），3），（（1,3），2），（（2,3），1），记为（1，2，3）
很显然，这形成一个递归过程，所以四个数字a，b，c，d的所有组合是（（a，b，c），d），（a，（b，c，d）），（（a，c，d），b），（（a，b，d），c）（（a，b），（c，d）），（（a，c），（b，d）），（（a，c），（b，d））
观察得知，这个递归过程可以归纳为任意数字组合后，与其他任意数，或其他任意数的组合，进一步组合

所以我们可以遍历数组里任意两个数字，把它们的所有组合存放在另一个数组里，，然后把原数组这两个数字删掉，把组合依次放进原数组里，进行递归过程，与原数组其他数字或其他数字的组合，递归组合成更高阶的组合，直到原数组只剩下一个数字，如果这个数字不是24，说明这一个组合不能计算出24，那么便递归返回，尝试其他组合

这里面还有一些需要注意的地方

 1. 由于除法可以有带小数的运算，所以我们需要根据精度来确认两个数是否相等
 2. 以及除法运算的时候除数不能为0
 3. 在使用vector的erase和insert方法的时候要注意先后顺序，因为这两个操作都会造成下标改变

## 2、代码实现

```C++
class Solution {
public:
    bool judgePoint24(vector<int>& nums) {
        vector<double> arr(nums.begin(), nums.end());
        
        return Recursion(arr);
        
    } 

    bool Recursion(vector<double>& nums) {
      if(nums.size() == 1) { 
         if(abs(nums[0] - 24) < 0.001) {
            return true;
         }
         return false;
      }
      for(int i = 0; i < nums.size(); i++) {
        for(int j = i + 1; j < nums.size(); j++) {
          double p = nums[i], q = nums[j];
          vector<double> ops{p + q, p - q, q - p, p * q};
          if(p > 0.001) ops.push_back(q / p);
          if(q > 0.001) ops.push_back(p / q);

          
          nums.erase(nums.begin() + j);
          nums.erase(nums.begin() + i);

          for(auto op : ops) {
            nums.push_back(op);
            if(Recursion(nums)) {
                return true;
            };
            nums.pop_back();
          }

          nums.insert(nums.begin() + i, p);
          nums.insert(nums.begin() + j, q);
        }
      }
        return false;
    }
};
```

