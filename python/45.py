import unittest
from typing import List


class Solution:
    def jump(self, nums: List[int]) -> int:
        dp = [0] * len(nums)
        for i in range(len(nums)):
            for j in range(1, min(len(nums) - i, nums[i] + 1)):
                dp[i + j] = dp[i] + 1 if dp[i +
                                            j] == 0 else min(dp[i + j], dp[i] + 1)
        return dp[-1]


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.jump([2, 3, 1, 1, 4]), 2)
