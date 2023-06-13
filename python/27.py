from typing import List


class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        pos = 0

        for i in range(len(nums)):
            num = nums[i]

            if num != val:
                nums[pos] = num
                pos += 1

        return pos
