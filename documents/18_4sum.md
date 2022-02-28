## 题意

给定一个数组 `nums` 和一个整数 `target`, 返回 nums 大小为 4 且和为 target 的子数组.

## 解答

首先对 nums 进行从小到大排序, 然后遍历 nums 到 i, 接下来就可以将问题退化成参数为 nums[i+1:len(nums)] 和 target-nums[i] 的 [15 题(3sum 问题)](https://leetcode150.xhu.me/15), 然后同样使用双指针逼近法获得和为 target-nums[i] 的子数组, 加上 nums[i] 存入到结果数组, 最后返回结果二维数组即可.