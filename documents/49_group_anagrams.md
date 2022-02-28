## 题意

给定一个字符串数组 `strs`, 返回一个字符串二维数组, 其中每一项是一组 [anagrams](https://en.wikipedia.org/wiki/Anagram), 即字符相同但是顺序不同的一组字符串.

## 解答

这道题的难点是怎么判断几个字符串互为 anagrams, 这里的解法是, 因为题目限制每个字符串只包含小写字母, 那么我们可以使用一个 `times [26]int` 数组来储存各个字母出现的次数, 这样每一组 anagrams 转换成这个 times 数组之后的值一定是相同的.

剩下就是将包含同样 times 数组的字符串放到结果二维数组的同一项里, 下面的代码使用了一个类似 [Trie 树](https://en.wikipedia.org/wiki/Trie) 的方法来处理.