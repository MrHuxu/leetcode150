# 1. Two Sum ![badge](https://img.shields.io/badge/-easy-green?style=flat-square)

[题目链接](https://leetcode.com/problems/two-sum)

## 题意

给定一个整数数组 nums 和一个目标整数 target, 找到 nums 中的两项和为 target 的并返回这两项的数组下标.

## 解答

从前往后遍历数组, 使用一个 map 来存储数组中一项的值和其下标的对应关系, 对于遍历到的一个下标 i:
1. target 与 nums[i] 的差不在 map 中, 将 nums[i] 和 i 存储到 map 中, 继续向后遍历;
2. target 与 nums[i] 的差已经在 map 中, 则可以获得差值的下标, 和 i 组成数组返回.

## 代码
