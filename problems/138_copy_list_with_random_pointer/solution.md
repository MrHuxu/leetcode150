## 题意

给定一个链表的头节点 `head`, 链表中每一个节点除了 next 指针外还有一个 random 指针, 复制整个链表包括 random 指针, 返回新的链表的头节点.

## 解答

思路和 [133 题](https://leetcode150.xhu.me/133) 一样, 使用一个 map 保存已经访问过的节点避免重复访问, 然后递归扫链表复制即可.