import unittest
from typing import List


class Solution:
    def climbStairs(self, n: int) -> int:
        if n <= 2:
            return n

        dp: List[int] = [0] * (n + 1)
        dp[0], dp[1], dp[2] = 0, 1, 2
        for i in range(3, n + 1):
            dp[i] = dp[i - 1] + dp[i - 2]

        return dp[n]


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.climbStairs(2), 2)
        self.assertEqual(solution.climbStairs(3), 3)
