from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if not nums:
            return 0

        pre = nums[0]
        pos = 1

        for i in range(1, len(nums)):
            num = nums[i]
            if num != pre:
                pre = num

                nums[pos] = num
                pos += 1

        return pos
