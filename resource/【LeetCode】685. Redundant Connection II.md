In this problem, a rooted tree is a directed graph such that, there is exactly one node (the root) for which all other nodes are descendants of this node, plus every node has exactly one parent, except for the root node which has no parents.

The given input is a directed graph that started as a rooted tree with N nodes (with distinct values 1, 2, ..., N), with one additional directed edge added. The added edge has two different vertices chosen from 1 to N, and was not an edge that already existed.

The resulting graph is given as a 2D-array of edges. Each element of edges is a pair [u, v] that represents a directed edge connecting nodes u and v, where u is a parent of child v.

Return an edge that can be removed so that the resulting graph is a rooted tree of N nodes. If there are multiple answers, return the answer that occurs last in the given 2D-array.

**Example 1:**
>**Input:** [[1,2], [1,3], [2,3]]
**Output:** [2,3]
**Explanation:** The given directed graph will be like this:
 &emsp;&nbsp; 1
&emsp; /&emsp;\\
&emsp;v&emsp; v
&emsp;2 -->3

**Example 2:**
>**Input:** [[1,2], [2,3], [3,4], [4,1], [1,5]]
**Output:** [4,1]
**Explanation:** The given directed graph will be like this:
5 <- 1 -> 2
&emsp;&emsp;^&emsp;&emsp;|
&emsp;&emsp;|&emsp;&emsp;v
&emsp;&emsp;4 <- 3

**Note:**
- The size of the input 2D-array will be between 3 and 1000.
- Every integer represented in the 2D-array will be between 1 and N, where N is the size of the input array.\

## 0、题目分析

题目中定义的树，是指有一个根节点，并且除了根节点外的其他节点都只有一个父节点的有向图

题目要求我们删掉给定的有向图的一条边，使得这个有向图成为一颗树

## 1、算法实现
那么最简单直接的思路就是每次删除一条边，对根节点进行一次DFS，如果能够遍历完所有节点，就说明这条边应该被删除，因为如果多了这条边，那么这棵树的某个节点就会多一个父节点，就不是题目要求的树了

问题在于如何找到根节点，我们可以计算每个节点的入度：

入度为0的节点就一定是根节点了；

如果没有入度为0的节点，那么所有节点的入度一定都为1；

因为一颗N个节点的树只有N-1条边，再加一条边的话，要么使根节点入度为1，要么使某个子节点入度为2，前一种情况则使所有节点入度为1

然后我们每次删除一条边，再对入度为0的节点进行一次dfs，没有入度为0的节点的时候，就对这条被删除的边指向的节点进行一次dfs，即若被删除的边为e：u-->v，对v进行dfs遍历，因为所有节点入度都为1，删除e之后，v的入度就变为0，也就成为了根节点

我们还可以把删除的边筛选一下，减少我们的算法复杂度，回到我们上面说的，一棵树再加一条边，要么使根节点入度为1，要么使某个子节点入度为2，前一种情况我们已经考虑完，后一种情况可以作为有根节点时边的筛选条件，即被删除的边指向的节点入度一定为2，举例来说就是e：u-->v，如果v的入度不为2，那么这条边就不能被删除，这样就省下了一次对根节点的dfs

## 2、代码实现

```c++
class Solution {
public:
    vector<int> findRedundantDirectedConnection(vector<vector<int>>& edges) {
      vector<vector<int>> graph(1001);
      vector<int> indegree(1001, 0);
      
      int nodeNum = 0;
      for(auto edge : edges) {
        graph[edge[0]].push_back(edge[1]);
        indegree[edge[1]]++;
        if(nodeNum < edge[0]) {
          nodeNum = edge[0];
        }
        if(nodeNum < edge[1]) {
          nodeNum = edge[1];
        }
      }

      int root = 0;
      for(int i = 1; i <= nodeNum; i++) {
        if(indegree[i] == 0) {
          root = i;
        }
      }
      for(int i = edges.size() - 1; i >= 0; i--) {
        vector<int> edge = edges[i];
        vector<bool> visited(1000, false);
          
        int num = 0;
        if(root != 0) {
          if(indegree[edge[1]] != 2) {
              continue;
          }
          dfs(graph, root, visited, num, edge[0], edge[1]);
        } else {
          dfs(graph, edge[1], visited, num, edge[0], edge[1]);
        }
        if(num == nodeNum) {
          return edge;
        }
        
      }
    }

    void dfs(vector<vector<int>> &graph, int node, vector<bool> &visited, int & num, int edge1, int edge2) {
      visited[node] = true;
      num++;
      for(auto i : graph[node]) {
        if(edge1 == node && i == edge2) {
            continue;
        }
        if(visited[i] == false) {
          dfs(graph, i, visited, num, edge1, edge2);
        }
      }
      return;
    }
};
```

