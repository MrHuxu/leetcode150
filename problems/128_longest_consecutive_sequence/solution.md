## 题意

给定一个整数数组 `nums`, 返回其中包含的最长连续整数序列的长度.

> [100, 4, 200, 1, 3, 2]  
  最长连续序列为 [1, 2, 3, 4], 返回 4

## 解答

我们使用一个 map `mapHeadToTail map[int]int` 来将一个连续序列的开头 map 到结尾, 另一个 `mapTailToHead map[int]int` 将一个连续序列的结尾 map 到开头.

假设数组里已经有了 `1,2,3` 这个序列, 那么两个 map 里面应该有 `mapHeadToTail[1]=3`, `mapTailToHead[3]=1`.

再使用一个 `existance map[int]bool` 来记录数组中元素的使用情况. 然后逐个遍历数组元素 `num`, 如果 existance[num] 为 true, 表示该数字已经使用过, 直接跳过这次循环, 否则把 existance[num] 置为 true. 然后通过 mapHeadToTail 拿到 num+1 对应的 `tail`, 通过 mapTailToHead 拿到 num-1 对应的 `head`:

1. **如果 head 和 tail 都存在**, 表示 num 是一个连续序列的中间项, 通过 num 可以把两个连续序列连起来, 则可以把 result 更新为 max(result, tail-head+1);
2. **如果只有 tail 存在**, 即 num 可以连到一个序列的开头, 则可以把 result 更新为 max(result, tail-num+1);
3. **如果只有 head 存在**, 即 num 可以连到一个序列的末尾, 则可以把 result 更新为 max(result, num-head+1);
4. **如果都不存在**, 则 result 为 max(result, 1).

然后在每一次更新完需要更新 mapHeadToTail 和 mapTailToHead 的值, 最后返回 result 即可.