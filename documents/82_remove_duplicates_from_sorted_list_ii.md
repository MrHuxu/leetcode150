## 题意

给定一个有序链表的头节点 `head`, 去掉链表中值重复的节点并返回最后的头节点.

## 解答

因为这道题的头节点也可能和别的节点重复, 所以需要创建一个 dummy 头节点出来.

具体的解法还是使用快慢指针法, 声明一个慢节点 slow=dummyHead, 一个快节点 fast=dummyHead.Next, 保持 fast 在 slow 前一个的位置, 开始遍历:

1. **如果 fast.Next.Val==slow.Next.Val**, 表示存在重复节点, 直接向后移动 fast, 并且把 fast.Next 接到 slow.Next 并后移 fast;
2. 如果不符合 1, 直接向后移动 slow 和 fast.

最后返回 dummy 头节点的 Next 指针即可.