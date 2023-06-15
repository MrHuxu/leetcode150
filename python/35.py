import unittest
from typing import List


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        if target <= nums[0]:
            return 0
        if target > nums[-1]:
            return len(nums)

        left, right = 0, len(nums) - 1
        while left < right:
            mid = (left + right) // 2
            if nums[mid] < target:
                left = mid + 1
            else:
                right = mid

        return (left + right) // 2


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.searchInsert([1, 3, 5, 6], 5), 2)
        self.assertEqual(solution.searchInsert([1, 3, 5, 6], 2), 1)
