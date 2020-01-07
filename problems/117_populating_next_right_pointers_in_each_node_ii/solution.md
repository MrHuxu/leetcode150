## 题意

给定一棵二叉树的根节点 `root`, 给每一个节点加上 next 指针, 使其指向当前节点在二叉树中右边的一个节点.

## 解答

这道题和 [116 题](https://leetcode150.xhu.me/116) 的区别就是这棵二叉树并不是一棵满二叉树, 左子节点不一定存在, 那么就不能直接沿着左子节点向下遍历.

这道题同样可以使用 dummy 头节点的方式来简化操作, 第一层为 dummyHead->root, 每遍历到一层, 就为下一层也建立一个头节点 nextHead, 并且从 dummyHead 使用 next 指针向右遍历, 把左子节点和右子节点连到 nextHead 的 next 指针上(若存在), 再把 dummyHead 更新为 nextHead. 如果遍历到 dummyHead 的 next 指针为空, 表示当前层没有节点了, 跳出循环即可.

最后返回原有根节点 root.