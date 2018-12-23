Given an undirected graph, return true if and only if it is bipartite.

Recall that a graph is bipartite if we can split it's set of nodes into two independent subsets A and B such that every edge in the graph has one node in A and another node in B.

The graph is given in the following form: graph[i] is a list of indexes j for which the edge between nodes i and j exists.  Each node is an integer between 0 and graph.length - 1.  There are no self edges or parallel edges: graph[i] does not contain i, and it doesn't contain any element twice.

>**Example 1:**
**Input:** [[1,3], [0,2], [1,3], [0,2]]
**Output:** true
**Explanation:** 
The graph looks like this:
0----1
|&emsp;&emsp;|
|&emsp;&emsp;|
3----2
We can divide the vertices into two groups: {0, 2} and {1, 3}.

.
>**Example 2:**
**Input:** [[1,2,3], [0,2], [0,1,3], [0,2]]
**Output:** false
**Explanation:** 
The graph looks like this:
0----1
| \ &emsp;|
|&emsp; \ |
3----2
We cannot find a way to divide the set of nodes into two independent subsets.

Note:

- graph will have length in range **[1, 100]**.
- graph[**i**] will contain integers in range **[0, graph.length - 1]**.
- graph[**i**] will not contain **i** or duplicate values.
- The graph is undirected: if any element **j** is in graph[**i**], then **i** will be in graph[**j**].

## 0、题目分析

题目要求判断一个无向图是不是二分图，那么什么是二分图呢？
>二分图又称作二部图，是图论中的一种特殊模型。 设G=(V,E)是一个无向图，如果顶点V可分割为两个互不相交的子集(A,B)，并且图中的每条边（i，j）所关联的两个顶点i和j分别属于这两个不同的顶点集(i in A,j in B)，则称图G为一个二分图。

## 1、算法实现
我们可以使用染色法来判定一个无向图不是二分图

我们将所有有边相连的两个邻接顶点分别用两种不同颜色染色，如果一个无向图是二分图，那么它最终可以被染成两种不同的颜色，因为任意三个点至多只有两条边相连，如果有超过两条边，就不是二分图了

那么如何给顶点们染色，我们使用DFS算法，给未被访问的节点**v**染上一种颜色，如何遍历相邻的节点**u**，给这些节点染上另一种颜色

显而易见，遍历的时候会出现三种情况
1. **u**已被访问且和**v**颜色不同
2. **u**已被访问且和**v**颜色相同
3. **u**未被访问

如果碰到第二种情况，这个图就不是二分图了

当所有节点都被访问且未遇到第二种情况，那这个图就是二分图

## 2、代码实现

C++实现
```C++
class Solution {
public:
    bool isBipartite(vector<vector<int>>& graph) {
        
        int nodeNum = graph.size();
        vector<int> visited(nodeNum);
        for(int i = 0; i < nodeNum; i++) {
        	if(visited[i] == 0 && !explore(graph, i, 1, visited)) {
        		return false;
        	}
        }
        return true;
    }

    bool explore(vector<vector<int>>& graph, int node, int color, vector<int>& visited) {
    	if(visited[node] != 0) {
    		return visited[node] == color;
    	}
    	visited[node] = color;
    	for(auto i : graph[node]) {
    		if(!explore(graph, i, -1 * color, visited)){
    			return false;
    		}
    	}
    	return true;
    }
};
```

go实现
```Go
func isBipartite(graph [][]int) bool {
    nodeNum := len(graph)
		visited := make([]int, nodeNum)
		for i := 0; i < nodeNum; i++ {
      if visited[i] == 0 && !explore(graph, i, 1, visited) {
        return false
      }
		}
		return true  
}

func explore(graph [][]int, node int, color int, visited []int) bool {
	if visited[node] != 0 {
		return visited[node] == color
	}
	visited[node] = color
	for i := 0; i < len(graph[node]); i++ {
		if !explore(graph, graph[node][i], -1 * color, visited) {
			return false
		}
	}
	return true
}
```
