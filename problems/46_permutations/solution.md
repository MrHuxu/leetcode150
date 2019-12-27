## 题意

给定一个整数数组 `nums`, 返回这个数组的排列 [Permutations](https://en.wikipedia.org/wiki/Permutation).

## 解答

从 Permutation 到定义可以看出来一个数组的排列就是把数组从 小->大 的排列通过把较大的数字不断往前移最后得到 大->小 的排列, 用 DFS 就可以搞定.