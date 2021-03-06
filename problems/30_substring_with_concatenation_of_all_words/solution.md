## 题意

给定一个字符串 `s` 和字符串数组 `words`, 找出 s 中所有包含 words 中所有字母的子串的起始位置.

## 解答

可以使用 DFS 方式求解.

使用一个 map 来储存每一个 word 出现的次数, 对遍历的每一个起始位置, 记下当前剩余的 word 数量, 以及把这个 map 传进去. 当找到一个 word 并对 map 里的数字进行 -1 操作时, 遍历完成后, 要对这个 word 进行 +1 操作, 也就是 `恢复现场`. 当遍历到剩余 word 数量为 0 时, dfs 函数返回 true 并将当前位置加入结果数组, 最后将结果数组返回.

**使用深度优先遍历解决问题的一个误区就是每个分支都拷贝一份状态(比如这题就是复制 map), 其实可以直接用同一个状态变量(比如这题的 map)来记录并且遍历完成恢复现场(即这题的 -1 和 +1 操作), 这样可以保证每一个分支的处理都不受前面遍历过的分支影响**
