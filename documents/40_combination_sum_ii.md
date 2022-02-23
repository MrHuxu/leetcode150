## 题意

给定一个整数数组 `candidates` 和整数 `target`, 返回 candidates 中所有和为 target 的组合,  和 [39 题](https://leetcode150.xhu.me/39) 不同的是 candidates 可能存在重复数字, 且每一个数字不可重复使用, 并且要求结果里不能出现重复数组.

## 解答

解题思路和 [39 题](https://leetcode150.xhu.me/39) 类似, 首先对 candidates 进行排序, 使其从小到大有序.

然后使用一个 combine 函数进行递归出来, 并且使用 `sum` 来记录当前数组的和, 和 `arr` 记录当前的数组, 每次递归:

1. **如果 sum 等于 target**, 那么把 arr 加入到结果数组;
2. **如果 sum > target**, 那么没必要再往后递归(这就是要先进行从小到大排序的原因), 直接返回;
3. **如果不满足 1 和 2**, 就在 sum 上依次加上 candidates 剩下的数字, 注意代码中 `i>idx && candidates[i]==candidates[i-1]` 这个条件, **这样可以保证每一个重复的数字, 只会在递归中遍历到, 而不会在这个 for 循环中重复遍历**, 同时更新 arr 往后递归.

最后返回结果数组即可.