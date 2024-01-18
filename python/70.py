import unittest
from typing import List


class Solution:
    def climbStairs(self, n: int) -> int:
        if n <= 2:
            return n

        dp1, dp2 = 1, 2
        for _ in range(3, n + 1):
            dp1, dp2 = dp2, dp1 + dp2
        return dp2


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.climbStairs(2), 2)
        self.assertEqual(solution.climbStairs(3), 3)
