## 题意

给定一个整数数组 `nums`, 返回这个数组的排列 [Permutations](https://en.wikipedia.org/wiki/Permutation), 和 [46 题](https://leetcode150.xhu.me/46) 不同的是, nums 里可能有重复数字, 需要对结果数组进行去重处理.

## 解答

首先我们改造一下 [31 题](https://leetcode150.xhu.me/31) 的 `nextPermutation` 函数, 使其返回一个 bool 值, 表示在这一次求解 next permutation 的过程有没有进行过位置的调换.

然后对 nums 进行从小到大排序, 然后对数组不断调用 nextPermutation 函数, 并且对返回的 bool 值进行判断, 如果返回为 false, 表示这一次没有进行数字调换, 已经得到了 大->小 排序的排列, 循环结束.