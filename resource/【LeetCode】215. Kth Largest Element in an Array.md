## 0、算法设计

考虑这样一种分治算法。对于任意给定的数v，假设S中的数被分成三组：比v小的数、与v相等的数（可能会有多个）以及比v大的数。分别记这三组数为S<sub>L</sub>、S<sub>V</sub>和S<sub>R</sub>。例如，如果集合S如下所示：
$$
S：2|36|5|21|8|13|11|20|5|4|1 
$$
	针对v=5，S被分为三组，分别为
$$S_L: 2|4|1$$
$$S_V: 5|5$$
$$S_R: 36|21|8|13|11|20$$
搜索范围立即缩小，转而在S的这三个子集中的某一个继续进行。如果我们想要寻找集合S的第8小元素，我们知道它一定是S<sub>R</sub>的第3小元素，因为|S<sub>L</sub>|+|S<sub>v</sub>|=5，即$$selection(S,8)=selection(S_R,3)$$
更一般的，通过比较k与这些子集所含元素数之间的大小关系，我们可以很快确定所要寻找的元素在哪个子集中：
$$
selection(S,k)=\left\{
\begin{array}{}
selection(S_L,k)& & {k \leq |S_L|}\\
v & & {|S_L| < k \leq |S_L|+|S_V|}\\
selection(S_L,k-|S_L|-|S_V|) & & { k > |S_L|+|S_V|}\\
\end{array} \right.
$$

## 1、算法的效率

如果每次选的v恰好是中项，则它的运行时间将满足
$$T(n)=T(n/2)+O(n)$$
即O(n)

但我们不可能每次选中的v正好是中项，我们可以采取一个简单的策略：从S中随机选取v

很显然，算法的运行时间依赖于我们随机选取的v。完全有可能出现这样的情况：由于持续的背运，我们挑选出来的v总是数组中最大的元素（或是最小元素）。这样的话，每次我们只能将待搜索的数组大小缩减1。对应的，在之前的例子里，我们可能首先选出的v=36，然后选出v=21，以此类推。这种最差情况将使得我们的选择问题算法必须要执行
$$n+(n-1)+(n-2)+...+\frac n2 = O(n^2)$$
次操作（当寻找中项时）。但是，这种情况出现的概率还是很低的。同样很少出现的还有最佳情形，即每次随机选取的v刚好能将数组一分为二，从而使算法的运行时间达到O(n)。在O(n)和O(n<sup>2</sup>)的范围内，算法的平均运行时间是多少呢？幸运的是，它与最佳情形下的运行时间很接近。

我们规定，当v落在它所在数组中的四分之一位置和四分之三位置之间，这样的v是好的。则一个随机选中的v较好的概率是50%，则在平均两次划分操作之后，待搜索数组的大小将最多缩减至原始大小的四分之三。

>基于这样的引理：在掷硬币游戏之中，在得到一次正面之前，平均需要掷一个匀质硬币2次。

则我们可以得到算法的期望时间：
$$T(n)\leq T( \frac {3n}4) + O(n) $$
由该递推式可得T(n)=O(n)
即算法的运行时间为O(n)

## 2、代码实现

```C++
int findKthLargest(vector<int>& nums, int k) {
	srand (time(NULL));
    vector<int> left, right;
    int leftCount = 0, middleCount = 0, rightCount = 0;
    int numSelected = nums[(rand() % nums.size())];
    for(auto num : nums) {
    	if(num < numSelected) {
    		left.push_back(num);
    		leftCount++;
    	} else if(num == numSelected) {
    		middleCount++;
    	} else {
    		right.push_back(num);
    		rightCount++;
    	}
    }
    if(k <= rightCount) {
    	return findKthLargest(right, k);
    } else if(k > rightCount && k <= rightCount + middleCount) {
    	return numSelected;
    } else {
    	return findKthLargest(left, k - rightCount - middleCount);
    }

}
```
