## 题意

给定一个有序链表集合 `lists`, 将集合里面的链表合并成一个新的有序链表, 并返回新的链表的头节点.

## 解答

方法1: 每次遍历从 lists 中所有链表的头结点找到最小值到节点并放到新的链表上, 直到所有链表都遍历完成, 每次找到最小值用堆实现的话, 在 LeetCode 上会超时.

方法2: 我们之前已经在 [21题](https://leetcode150.xhu.me/21) 中实现了合并两个有序链表的函数, 那么起始我们可以用归并的思想, 对 lists 中的链表进行两两归并, 直到最终 lists 里只剩下一个链表, 返回这个链表的头结点即可. 其中每次归并使用的 merge 函数可以直接使用 21 题的现成代码.