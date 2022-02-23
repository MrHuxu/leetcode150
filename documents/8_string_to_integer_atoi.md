## 题意

编写一个 `atoi` 函数. [C++ 的 atoi 函数定义](http://www.cplusplus.com/reference/cstdlib/atoi/)

## 解答

遍历每一位的字符和 `'0'` 求差即可求解, 需要注意几点:

1. 起始的空格需要忽略;
2. 起始的负号需要处理;
3. 返回的数字需要在 int32 的范围内.