## 题意

给定一个整数数组 `digits` 表示一个大数字的从高到低的各个位, 返回其加 1 的结果, 同样使用整数数组表示.

## 解答

使用一个 `plus` 变量保存进位情况, 逐位遍历, 最后根据进位情况补上 1, 返回即可.