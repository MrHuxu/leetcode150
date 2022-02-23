## 题意

给定一个整数数组 `prices` 表示某股票的每一天的价格, 求买卖一次股票可能获得的最大利润.

## 解答

从前向后遍历 prices, 用一个变量 `min` 记录当前位置之前最小的值:

1. 如果当前值大于 min, 则用差值更新结果;
2. 如果当前值小于等于 min, 则更新 min.

最后返回结果即可.