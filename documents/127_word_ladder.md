## 题意

给定两个字符串 `beginWord` 和 `endWord` 表示起始单词和结束单词, 一个字符串数组 `wordList` 表示可以使用的单词, beginWord 每次改变一个字母得到一个 wordList 单词, 最后变成 endWord, 判断这个转换最少有几个单词.

> beginWord = "hit"  
  endWord = "cog"  
  wordList = ["hot","dot","dog","lot","log","cog"]  
  "hit" -> "hot" -> "dot" -> "dog" -> "cog" 返回 5

## 解答

这道题可以使用 BFS 求解, 从 beginWord 开始每次收集当前的单词可以转换得到的所以单词, 再继续往下遍历, 一旦出现了 endWord, 就返回当前转换次数.

首先我们对 wordList 进行一下预处理, 生成一个 `changeMap`, 其内容为如下格式:

    map[ "*ot": []string{ "hot", "lot" } ] 

这样当我们需要找一个单词可以转换得到的单词列表时, 只需要依次将其各位置成 `*` 并作为 key, 合并 changeMap 里所有的 value 即可.

bfs 中注意使用一个 `used map[string]bool` 记录下已经使用过的单词进行剪枝, 具体过程就不详细描述了, 直接看代码.