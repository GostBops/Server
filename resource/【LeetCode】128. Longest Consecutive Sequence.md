
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
Your algorithm should run in O(n) complexity.

**Example:**
>**Input:** [100, 4, 200, 1, 3, 2]
**Output:** 4
**Explanation:** The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.


## 0 、题目分析
我们可以使用unordered_set和map两种STL来获取最长连续序列长度

## 1、map
1. 首先我们定义一个空map
2. 接着我们遍历每一个数字num
	1. 如果这个数字num不在map里面，即map[num]==0，我们就开始遍历它的左右，即num-1，num+1
	2. 如果num-1在map里面，即LeftValue = map[num-1]>0，如果不在，我们取LeftValue=0
	3. 同理取RightValue
	4. 将map[num]，map[num - LeftValue]和map[num - RightValue]赋值为LeftValue + RightValue + 1
	5. 将res更新为max（res，LeftValue + RightValue + 1）
	6. 如果这个数字num在map里面，说明它已经被遍历过，因此不做任何操作

```c
class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        int res = 0;
        map<int, int> m;
        for (int d : nums) {
            if (!m.count(d)) {
                int left = m.count(d - 1) ? m[d - 1] : 0;
                int right = m.count(d + 1) ? m[d + 1] : 0;
                int sum = left + right + 1;
                m[d] = sum;
                res = max(res, sum);
                m[d - left] = sum;
                m[d + right] = sum;
            }
        }
        return res;
    }
};
```

## 2、set
1. 首先把nums储存到unordered_set里面
2. 然后遍历每一个数字num
	1. 如果num在set里面，那么遍历num-1和num+2，直到num-left-1和num+right+1不在set里面
	2. 将res更新为max（res，right+left+1）
	3. 每一次查找成功时，需要把这个数字从set中删除，因为它已经不需要被查找，可以减少查找的复杂度

```c
class Solution {
public:
    int longestConsecutive(vector<int>& nums) {
        int res = 0;
        unordered_set<int> s(nums.begin(), nums.end());
        for (int val : nums) {
            if (!s.count(val)) continue;
            s.erase(val);
            int pre = val - 1, next = val + 1;
            while (s.count(pre)) s.erase(pre--);
            while (s.count(next)) s.erase(next++);
            res = max(res, next - pre - 1);
        }
        return res;
    }
};
```
