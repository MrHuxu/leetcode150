## 题意

给定一个有序链表的头节点 `head`, 删除链表里值重复的节点使同一个值在链表里只出现一次.

## 解答

因为这道题头节点最后肯定是会保留的, 所以不需要使用 dummy 头节点. 直接使快慢指针 `fast` 和 `slow` 都指向 head, fast 每次遍历到不等于 slow 的位置, 再把 slow.Next=fast, 后移 slow, 最后仍然返回 head 即可.