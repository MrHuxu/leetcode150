## 题意

给定一个整数数组 `nums`, 找到 nums 中的三项和为 0 的组合并返回这三项的数组下标.

## 解答

首先对数组进行排序.

然后对数组进行遍历, 当遍历到 i 时, 声明两个指针 `left=i+1` 和 `right=len(nums)-1`, 并且根据 `nums[i]+nums[left]+nums[right]` 和 0 的大小关系让 left 和 right 向中间逼近, 最后找到所有结果返回即可.