## 题意

给定一个数字 `x`, 返回其翻转的数字.

## 解答

直接使用递归求解, 每次将已有结果乘 10 并加上 x 的个位数字, 然后 x 进行除 10 进行递归, 当 x 为 0 的时候返回结果.

递归的时候需要判断数字, 如果已经超过了 int32 的表示范围, 直接返回 0 即可.