import unittest
from typing import List


class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        n, i = len(nums), 0
        while i < n:
            num = nums[i]
            if num > 0 and num <= n and nums[num - 1] != num:
                nums[num - 1], nums[i] = nums[i], nums[num - 1]
                i -= 1
            else:
                i += 1
        for i, num in enumerate(nums):
            if i + 1 != num:
                return i + 1
        return n + 1


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.firstMissingPositive([1, 2, 0]), 3)
        self.assertEqual(solution.firstMissingPositive([3, 4, -1, 1]), 2)
