## 题意

给定一个已经排好序的数组 `nums`, 清理里面的重复数据使同一个数字最多出现两次并且位置在数组的前部, 返回最后数组前面合法数据的长度, 这个长度之后的数据不影响题目结果.

> [1,1,1,2,2,3] 返回 5 且前 5 项为 [1,1,2,2,3]  
  [0,0,1,1,1,1,2,3,3] 返回 7 且前 7 项为 [0,0,1,1,2,3,3]

## 解答

首先我们使用一个 `preIndex` 变量来表示前面存放合法数据的下标并置初始值为 0, 使用一个布尔值 `same` 表示 preIndex 和 preIndex-1 的项是否相等, 再遍历 nums 获得下标 `i`:

1. **如果 nums[i] 等于 nums[preIndex]**, 那么这时就看 same 的值,
  - 如果 same 为 false 的话, preIndex 还可以向后移一个, 这时就把 preIndex 向后移一位, 并且将 nums[i] 的值赋给 nums[preIndex], 并把 same 值置为 true; 
  - 如果 same 为 true, 表示 preIndex 前已经有相同的值了, 不用处理.
2. **如果 nums[i] 不等于 nums[preIndex]**, 那么直接把 preIndex 后移, 并置 nums[preIndex]=nums[i], 再把 same 置为 false.

最后返回 preIndex+1 即可.