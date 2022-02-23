## 题意

给定三个字符串 `s1`, `s2` 和 `s3`, 判断 s3 是否是有 s1 和 s2 的字符互相穿插得到的.

## 解答

这道题可以使用 DP 求解.

首先如果 s1 和 s2 的长度和不等 s3 的长度的话, 直接返回 false 即可. 然后声明一个二维数组 `dp [][]bool` 使用 dp[i][j] 表示 s1 的前 i 个字符和 s2 的前 j 个字符互相穿插能否得到 s3 的前 i+j 个字符, 那么有如下的转换关系:

1. **如果 i 和 j 为 0**, 那么 dp[i][j]=true;
2. **如果只有 i 为 0**, 那么直接对比 s2 和 s3 的前 j 个字符即可;
3. **如果只有 j 为 0**, 那么直接对比 s1 和 s3 的前 i 个字符即可;
4. **如果 i 和 j 都不为 0**, 那么当 s3 的第 i+j 个字符等于 s2 的第 j 个字符时, dp[i][j] 可以通过 dp[i][j-1] 得到; 当 s3 的第 i+j 个字符等于 s1 的第 i 个字符时, dp[i][j] 可以通过 dp[i-1][j] 得到. 状态转移关系为: `dp[i][j]=(s2[j-1]==s3[i+j-1] && dp[i][j-1]) || (s1[i-1]==s3[i+j-1] && dp[i-1][j])`.

最后返回 dp[len(s1)][len(s2)] 即可.