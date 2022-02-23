## 题意

给定一个数字 `x`, 实现 [sqrt 函数](http://www.cplusplus.com/reference/cmath/sqrt/), 即返回这个数字的平方根.

## 解答

首先如果 x 为 0 或 1 的话, 直接返回 x 即可.

然后从 1 和 x/2 为左右边界进行二分查找:

1. 如果二分的中间值 mid 的平方等于 x, 直接返回 mid 即可;
2. 如果 x 第平方根不为整数, 那么在二分逼近 x 的过程中, 最后必定存在 right-1 的操作, 所以最后返回 right 即可.