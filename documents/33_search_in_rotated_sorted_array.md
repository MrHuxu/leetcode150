## 题意

给定一个数组 `nums` 和整数 `target`, 找到 target 中 nums 中的位置下标, 如果不存在的话返回 -1. nums 是一个被旋转过的(rotated)的整数数组.

> rotated: [0,1,2,4,5,6,7] -> [4,5,6,7,0,1,2]

## 解答

首先遍历数组, 通过判断 `nums[i]>nums[i+1]` 找到两段有序数组的分界位置, 再通过 target 和 nums[0] 的大小关系来确定 target 在这两段里的哪一部分, 再通过二分查找即可.