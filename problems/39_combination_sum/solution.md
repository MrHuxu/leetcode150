## 题意

给定一个整数数组 `candidates` 和整数 `target`, 返回 candidates 中所有和为 target 的组合, candidates 中不存在重复数字, 且每一个数字都可以重复使用.

## 解答

首先对 candidates 进行排序, 使其从小到大有序.

然后使用一个 combine 函数进行递归出来, 并且使用 `sum` 来记录当前数组的和, 和 `arr` 记录当前的数组, 每次递归:

1. **如果 sum 等于 target**, 那么把 arr 加入到结果数组;
2. **如果 sum > target**, 那么没必要再往后递归(这就是要先进行从小到大排序的原因), 直接返回;
3. **如果不满足 1 和 2**, 就在 sum 上依次加上 candidates 剩下的数字, 同时更新 arr 往后递归.

最后返回结果数组即可.