## 题意

给定一个二维数组 `obstacleGrid` 表示一个棋盘, 每一步只能从当前位置向右或向下移动, 遇到 1 的位置则不能通过, 初始位置在左上角, 结束位置在右下角, 返回总共的线路数, 

## 解答

解法和 [62 题](https://leetcode150.xhu.me/62) 一致, 在遍历到时候遇到 1 就忽略即可.