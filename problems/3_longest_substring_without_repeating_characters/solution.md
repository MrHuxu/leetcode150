## 题意

给定一个字符串 `s`, 找出里面最长的不含重复字符的子串的长度.

## 解答

使用滑动窗口的方法求解.

使用一个 front 变量记录窗口的起始位置, 以及一个 map 保存窗口里各个字符出现的次数, 当遍历到 i 时:

1. s[i] 在 map 中值为 0, 即当前窗口没有重复字符, 使用 `i-front+1` 来更新结果的值;
2. s[i] 在 map 中值不为 0, 表明这个字符已经出现过, 则想后移动 front 并同时将 map 中 s[front] 的值置为 0, 直到 s[front]==s[i] 结束, 然后将 front 往后移一位.

然后将 map 中当前字符的值置为 1, 继续遍历到结束, 最后返回结果即可.