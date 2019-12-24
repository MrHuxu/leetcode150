## 题意

给定一个整数数组 `nums` 和一个目标整数 `target`, 找到 nums 中的三项和最接近 target 的值并返回.

## 解答

解题思路和 [15题(3Sum)](https://leetcode150.xhu.me/15) 一致, 设置一个变量 `minSub` 来存储已经找到的三项和与 target 的差的最小值, 如果出现差值小于 minSub 的情况, 更新返回值即可.