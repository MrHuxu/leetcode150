## 题意

给定一个由数字组成的字符串 `s`, 已知如下的映射方式:

> 'A' -> 1
  'B' -> 2
  ...
  'Z' -> 26

返回 s 有多少种可以被转换成字母序列的方式.

> s="12", 因为可以转换成 "AB"(1 2) 也可以转换成 "L"(12), 返回 2.

## 解答

这道题可以使用 DP 求解, 声明一个数组 `dp []int` 使用 dp[i] 表示 s[0:i+1] 有多少种转换方式, 首先可以直接得到 dp[0] 和 dp[1] 即前 1 个和前 2 个字符有多少种方式.

对于 i>=2, 如果 s[i]!='0', 则 s[i] 可以独立转换成一个字母, 即 dp[i]+=dp[i-1], 如果 s[i-1:i+1] 这个字符串在 '10' 和 '26' 之间, 则 s[i-1:i+1] 可以转换成一个字母, 即 dp[i]+=dp[i-2].

最后返回 dp[len(s)-1] 即可.