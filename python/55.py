import unittest
from typing import List


class Solution:
    def canJump(self, nums: List[int]) -> bool:
        distance = 0
        for i, num in enumerate(nums):
            if distance >= len(nums) - 1:
                return True
            if num == 0 and distance <= i:
                return False
            distance = max(distance, i + num)
        return False


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(solution.canJump([2, 3, 1, 1, 4]))
