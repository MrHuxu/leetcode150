## 题意

给定一个矩阵 `matrix` 和一个整数 `target`, 这个矩阵有以下两个特征:

1. 每一行从左向右递增;
2. 每行第一项大于前一行最后一个项.

判断 target 是否在 matrix 中.

## 解答

两次二分查找, 第一次通过二分找到 target 所在的行 `targetRow`, 第二次通过二分判断 targetRow 中是否包含 target.