## 题意

将十进制数字转成罗马数字.

## 解答

可以用一个常量来储存各个位上的罗马数字, 然后对十进制数字逐位遍历, 再把数字分成以下的范围进行处理:

1. 大于 0 小于等于 3
2. 等于 4
3. 大于 4 小于等于 8
4. 等于 9