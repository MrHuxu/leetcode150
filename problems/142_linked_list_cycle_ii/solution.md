## 题意

给定一个链表的头节点 `head`, 找出这个链表里环的入口.

## 解答

还是使用快慢指针的方式求解:

1. 首先从头节点开始, 快指针一次遍历前进两个, 慢指针一次前进一个, 这样相遇的话可以找到环里面的一个节点 `nodeInList`, 找不到的话就直接返回 nil;
2. 再从 nodeInList 出发循环一圈可以获得环的长度 `listLen`;
3. 再使用快慢指针法, 快指针比慢指针先出发 listLen 个位置, 同时开始每次前进一个节点开始遍历, 相遇的位置就是环的入口.