## 题意

给定一个链表的头节点 `head`, 使用 [插入排序](https://en.wikipedia.org/wiki/Insertion_sort) 将其排序, 并返回排好序的链表的头节点.

## 解答

使用 dummy 头节点, 遍历 head 找到合适的位置插入到 dummy 节点后, 最后返回 dummytou 头节点的 next 指针即可.