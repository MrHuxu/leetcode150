## 题意

给定一个整数数组 `nums`, 返回这个数组的排列 [Permutations](https://en.wikipedia.org/wiki/Permutation), 和 [46 题](https://leetcode150.xhu.me/46) 不同的是, nums 里可能有重复数字, 需要对结果数组进行去重处理.

## 解答

这题同样可以使用 DFS 求解, 关键在于 `i > 0 && nums[i] == nums[i-1] && !used[i-1]` 这个条件, 当 `used[i-1]=true` 时, 表示相邻的 nums[i-1] 和 nums[i] 在一个遍历分支里, 需要执行下去, 当 `used[i-1]=false` 时, 表示跳过 nums[i-1] 对相同的 nums[i] 执行了相同的分支, 需要被剪枝, 这样遍历完可得结果.