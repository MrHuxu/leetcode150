## 题意

给定一个有序数组 `nums`, 使用数组中元素构造一棵二叉搜索树并返回其根节点, 要求这棵二叉树尽可能的浅, 尽可能贴近平衡二叉树.

## 解答

题目要求这棵二叉树尽可能平衡, 那么就必须保证左右子树的节点数量尽可能一致.

我们可以找到这个数组的中点 `mid` 并作为根节点, 那么左子树遍历结果为 nums[0:mid], 右子树遍历结果为 nums[mid+1:len(nums)], 递归构造左右子树即可.