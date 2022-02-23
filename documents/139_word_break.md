## 题意

给定一个字符串 `s` 和数组 `wordDict`, 判断 s 是否可以分割成若干个 wordDict 中的单词, 即 s 是否全是有 wordDict 中的单词组成.

## 解答

首先为了方便判断一个单词是否在 wordDict 中, 我们使用一个 map `exist map[string]bool`, 然后声明一个数组 `dp []bool` 使用 dp[i] 表示 s[0:i] 是否是由 wordDict 中单词构成.

然后遍历数组获得 i, 再从 i 向前遍历到 j, 如果 exist[s[j:i]] 即当前这个子串在 wordDict 中, 且 dp[j] 即前 j 个字符是有 wordDict 中单词构成, 那么 dp[i]=true. 然后继续向后遍历.

实际遍历中, 为了减少比较次数, 因为 wordDict 里单词的长度有一个范围 `minLen<=len<=maxLen`, 不需要每次都从 i 遍历到 0, 只需要遍历单词的长度范围即可.

最后返回 dp[len(s)].