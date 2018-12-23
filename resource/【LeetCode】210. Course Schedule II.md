There are a total of n courses you have to take, labeled from 0 to n-1.

Some courses may have prerequisites, for example to take course 0 you have to first take course 1, which is expressed as a pair: [0,1]

Given the total number of courses and a list of prerequisite pairs, return the ordering of courses you should take to finish all courses.

There may be multiple correct orders, you just need to return one of them. If it is impossible to finish all courses, return an empty array.

**Example 1:**
>**Input**: 2, [[1,0]] 
**Output**: [0,1]
**Explanation**: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].

**Example 2:**
>**Input**: 4, [[1,0],[2,0],[3,1],[3,2]]
**Output**: [0,1,2,3] or [0,2,1,3]
**Explanation**: There are a total of 4 courses to take. To take course 3 you should have finished both courses 1 and 2. Both courses 1 and 2 should be taken after you finished course 0. So one correct course order is [0,1,2,3]. Another correct ordering is [0,2,1,3] .

**Note**:

 1. The input prerequisites is a graph represented by a list of edges, not adjacency matrices. Read more about how a graph is represented.
 2. You may assume that there are no duplicate edges in the input prerequisites.
 

## 0、题目分析

显而易见，这是一道拓扑排序的题目
>对一个有向无环图(Directed Acyclic Graph简称DAG)G进行拓扑排序，是将G中所有顶点排成一个线性序列，使得图中任意一对顶点u和v，若边(u,v)∈E(G)，则u在线性序列中出现在v之前。 通常，这样的线性序列称为满足拓扑次序(Topological Order)的序列，简称拓扑序列。

## 1、算法设计

那么拓扑排序应该怎么实现呢？
很简单，对有向无环图G进行一次深度优先搜索（DFS）
>**深度优先搜索算法（英语：Depth-First-Search，DFS**）是一种用于遍历或搜索树或图的算法。沿着树的深度遍历树的节点，尽可能深的搜索树的分支。当节点v的所在边都己被探寻过，搜索将回溯到发现节点v的那条边的起始节点。这一过程一直进行到已发现从源节点可达的所有节点为止。如果还存在未被发现的节点，则选择其中一个作为源节点并重复以上过程，整个进程反复进行直到所有节点都被访问为止。属于盲目搜索。
```C++
//无向图的深度优先搜索
//Input: G=(V,E) is a graph; v∈V
//Output: visited(u) is set to true for all nodes u reachable from v

procedure explore(G,v)
    visited(v) = true
    previsit(v)
    for each edge (v,u) ∈ E:
        if not visited(u): explore(u)
    postvisit(v)

procedure dfs(G)
    for all v ∈ V:
        visited(v) = false
    for all v ∈ V:
        if not visited(v): explore(v)
```

我们将在图的遍历过程中多搜集一些信息：对于每个顶点，我们将记录下两个重要事件发生的时间，一个是每个顶点最先被访问（previsit）的时刻，另一个是最后离开每个顶点（postvisit）的时刻
```C++
procedure previsit(v)
    pre[v] = clock
    clock = clock + 1
    
procedure postvisit(v)
    post[v] = clock
    clock = clock + 1
```

对于有向图来说，如果将post从大到小排列，则会形成一个可行的拓扑排序，因为后代节点肯定比它的祖先节点更早结束访问，即post<sub>后代</sub> < post<sub>祖先</sub>,只要祖先节点排在后代节点的前面，就能形成拓扑排序

那么我们的算法就基本完成了：
```C++
procedure dfs(G)
    for all v ∈ V:
        visited(v) = false
    for all v ∈ V:
        if not visited(v): explore(v)
    sort(post)
    
procedure explore(G,v)
    visited(v) = true
    previsit(v)
    for each edge (v,u) ∈ E:
        if not visited(u): explore(u)
    postvisit(v)
    
procedure previsit(v)
    pre[v] = clock
    clock = clock + 1
    
procedure postvisit(v)
    post[v] = clock
    clock = clock + 1
```

不过我们还可以改进一下
比如我们并不需要储存pre，而且我们不关心每个顶点最先被访问的时刻
实际上我们也并不需要排序，我们只需要在每个节点结束访问时把它推到栈中就行，因此post也是不需要存储的

```c++
procedure dfs(G)
    for all v ∈ V:
        visited(v) = false
    for all v ∈ V:
        if not visited(v): explore(v)
    return sort
    
procedure explore(G,v)
    visited(v) = true
    for each edge (v,u) ∈ E:
        if not visited(u): explore(u)
    postvisit(v) = true
    sort: push v into Stack
```

不过我们忘了一个大前提，输入的图并一定是有向无环图，也可能是有环的，当有向图有环时，无法形成一个拓扑排序。那么我们怎么判断有向图是否有环呢？

我们需要定义一种边的类型，回边：
>回边是DFS树中从一个顶点指向其祖先的边

如何判断一条边是不是回边：
>对于一条边（u，v），若当u遍历到这条边时发现v已被访问却未结束访问，则说明v是u的祖先，即这条边是回边
```
visited[i] == true && post[i] == false//回边
```

>有向图含有一个环当且仅当深度优先搜索过程中探测到一条回边

所以当存在回边时，有向图有环，不能形成拓扑排序

## 2、代码实现

```C++
class Solution {
public:
	int numSort;
    vector<int> findOrder(int numCourses, vector<pair<int, int>>& prerequisites) {
        numSort = numCourses;
        vector<vector<int>> graph(numCourses, vector<int>(0));
        for(auto edge : prerequisites) {
        	graph[edge.second].push_back(edge.first);
        }
        
        vector<bool> post(numCourses);
        vector<bool> visited(numCourses);
        vector<int> sort(numCourses);
        int i;
        for(i = 0; i < numCourses; i++) {
        	if(visited[i] == false && !dfs(graph, i, visited, post, sort)) {
        		vector<int> v;
                return v;
        	}
        }

        return sort; 
    }

    bool dfs(vector<vector<int>>& graph, int node, vector<bool>& visited, vector<bool>& post, vector<int>& sort) {
    	visited[node] = true;
        for(auto i : graph[node]) {
            if(visited[i] == true && post[i] == false) {
                return false;
            }
            if(visited[i] == false && !dfs(graph, i, visited, post, sort)){
                return false;
            }
        }
    	post[node] = true;
    	numSort--;
    	sort[numSort] = node;
    	return true;
    }
};
```


