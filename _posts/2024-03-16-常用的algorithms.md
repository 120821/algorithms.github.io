---
layout: post
title: "常用的algorithms"
date: 2024-03-16
categories: [go algorithms]
---
常用算法主要有以下几类

### 1. 排序算法
- **快速排序**：通过选定一个基准值，将数组分为小于和大于基准值的两部分，递归排序子数组。
- **归并排序**：将数组分成最小单位，然后两两合并时进行排序，直至整个数组变为有序。
- **堆排序**：利用堆这种数据结构所设计的一种排序算法，通过构建最大堆或最小堆，实现数组的排序。

### 2. 搜索算法
- **二分查找**：在有序数组中，通过不断将搜索区间减半来寻找特定值。
- **深度优先搜索（DFS）**：采用递归方法，尝试每一条路径直到找到解为止。
- **广度优先搜索（BFS）**：从根节点开始，逐层向外扩展，直到找到解为止。

### 3. 动态规划
- **背包问题**：给定一组物品，每个物品都有重量和价值，确定在限定的总重量内，如何选择物品以使得总价值最大。
- **最长公共子序列（LCS）**：找出两个序列共有的、顺序一致但不必连续的最长子序列。
- **最短路径问题**：如Dijkstra算法，用于在加权图中找到一个顶点到其他顶点的最短路径。

### 4. 图算法
- **图的遍历**：如深度优先搜索（DFS）和广度优先搜索（BFS）。
- **最短路径**：如Dijkstra算法和Bellman-Ford算法。
- **最小生成树**：如Prim算法和Kruskal算法，用于找到图的最小生成树，保证所有节点以最小的连接成本连接。

### 5. 数据结构相关算法
- **二叉树遍历**：包括前序、中序、后序遍历，可以递归或迭代实现。
- **图的表示和遍历**：邻接矩阵和邻接表的表示，以及基于它们的DFS和BFS遍历。
- **哈希算法**：设计良好的哈希函数，以及处理哈希冲突的策略，如链地址法（拉链法）和开放地址法（线性探测、二次探测）。

### 6. 其他算法
- **字符串匹配算法**：如KMP算法，用于在主文本中找到一个匹配子串的高效算法。
- **贪心算法**：在每一步选择中都采取在当前状态下最好或最优（即最有利）的选择，以期望导致全局最好或最优的解。
- **分治算法**：将原问题划分成n个规模较小而结构与原问题相似的子问题，递归解决这些子问题，然后再合并其结果，得到原问题的解。

了解并熟练这些算法的思想和实现非常重要。

了解每种算法的时间复杂度和空间复杂度，以及它们的优缺点和适用场景。

