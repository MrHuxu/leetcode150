## 题意

给定一个整数 `n`, 返回二进制长度为 n 的一组格雷码.

> A Gray code is an encoding of numbers so that adjacent numbers have a single digit differing by 1.

也就是相邻两项的二进制表示只有一位的差别.

## 解答

首先对于 n=1, 我们很容易得到这样的格雷码 `[0, 1]`.

然后当 n=2 的时候, 我们对 n=1 的结果进行倒序遍历, 再在最高位加上 1 加入到结果中, 那么可以得到 `[00, 01, 11, 10]`, 是一个合法的格雷码序列.

n=3 时同样这么处理, 可以得到 `[000, 001, 011, 010, 110, 111, 101, 100]`, 同样合法.

那么我们可以得到规律就是, 对于二进制位数为 n 的格雷码序列, 可以通过对 n-1 的序列进行逆序遍历再在最高位置 1 并依次加到原序列尾部得到.