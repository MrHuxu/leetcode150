## 题意

给定一个链表的头节点 `head`, 判断这个链表里是否有环.

## 解答

可以使用快慢指针的方法, 一个指针每次前进两步, 一个每次前进一步, 如果两个指针相遇, 则存在环, 如果快指针遇到了 nil, 则不存在环.