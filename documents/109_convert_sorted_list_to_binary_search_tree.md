## 题意

给定一个有序链表的头节点 `head`, 使用链表中元素构造一棵二叉搜索树并返回其根节点, 要求这棵二叉树尽可能的浅, 尽可能贴近平衡二叉树. 

## 解答

首先获得链表的长度, 然后首先对左边这一半长度的链表递归构建左子树, 同时改变 `head` 指针的内容, 那么当左子树构建完成时, head 指向的是链表的下一个节点, 也就是左右子树的中间位置, 将这个位置作为根节点, 对剩下的链表递归构建右子树即可.