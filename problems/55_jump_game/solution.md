## 题意

给定一个整数数组 `nums`, nums[i] 表示在位置 i 可以向后跳的距离, 判断从 nums[0] 开始能否跳到数组末尾.

## 解答

遍历数组, 并且使用一个变量 `distance` 表示当前可以跳到的最远的距离, 每次遍历都把 distance 更新为 distance 和 nums[i]+i 的较大值, 那么每次遍历有下面两种情况:

1. **distance >= len(nums)-1**, 即可以达到数组末尾, 这时直接返回 true 即可;
2. **distance <= i 且 nums[i] == 0**, 即 distance 达不到当前的位置 i 且当前位置无法向后跳, 这时肯定跳不过 i 这个位置, 直接返回 false.

如果循环结束仍然没有返回的话, 就直接返回 false 即可.