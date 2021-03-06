## 题意

给定一个二维数组 `triangle` 表示一个等边三角形, 从顶点开始向下移动, 每一次只能移动到下一层相邻的点, 求从顶点到底边和最短的路径.

## 解答

我们直接从上往下更新这个二维数组, 对于每一层来说, 其第一个和最后一个点只能由上一层第一个和最后一个点移动得到, 而对于中间的点, 从上一层相邻的两个点取更小的那一个加上当前点的值就是到当前点的最小路径, 即 triangle[i][j]+=min(triangle[i-1][j-1],triangle[i-1][j]).

最后返回 triangle 最底层的最小值即可.