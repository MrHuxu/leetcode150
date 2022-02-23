## 题意

给定一个字符串 `s`, 求最小的分割次数, 可以把这个字符串分割成若干个回文子串.

## 解答

首先声明二维数组 `dp [][]bool` 使用 dp[i][j] 表示 s[i:j+1] 是否为回文子串, dp 的值可以通过和 [131 题](https://leetcode150.xhu.me/131) 一样的方法得到.

再声明一个数组 `dp2 []int` 使用 dp[i] 表示 s[0:i+1] 需要的最小分割次数, 遍历 i 从 1 到 len(s)-1 (为什么不需要 0? 因为 s[0:1] 只有一个字符直接就是回文了, 分隔次数为 0), 首先置 dp2[i]=i, 即最差情况为完全没有回文子串, 每个字符间都要分割一次. 然后对于任意 0<=j<=i, 如果 dp[j][i] 为 true:

1. 如果 j 为 0, 那么表示 s[0:i+1] 为回文子串, 不需要分割, dp2[i]=0;
2. 如果 j 不为 0, 那么 dp2[i] 可以表示为 dp2[j-1] 后面接一个 s[j:i+1] 这个回文子串, 可以得到转移方程 dp2[i]=min(dp2[i], dp2[j-1]+1).

最后返回 dp2[len(s)-1] 即可.