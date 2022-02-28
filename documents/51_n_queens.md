## 题意

给定一个正整数 `n`, 返回 n 皇后问题的所有的解.

> 通过 [八皇后问题](https://en.wikipedia.org/wiki/Eight_queens_puzzle) 扩展的 n 皇后问题, 即在一个 n*n 的棋盘上, 摆放 n 个皇后子, 其中任意两个皇后不在同一条 横/竖/斜 线上.

## 解答

算法原理和 [52 题](https://leetcode150.xhu.me/52) 一样, 是需要额外使用一个二位数 `board [][]byte` 来保存棋盘内容, 当出现合法解时复制一份存到结果返回即可.