import unittest
from typing import List


class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        if not grid:
            return 0

        dp = [[0] * len(grid[0]) for _ in range(len(grid))]
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                dp[i][j] = grid[i][j]

                if i >= 1 and j >= 1:
                    dp[i][j] += min(dp[i - 1][j], dp[i][j - 1])
                elif i >= 1:
                    dp[i][j] += dp[i - 1][j]
                elif j >= 1:
                    dp[i][j] += dp[i][j - 1]

        return dp[len(grid) - 1][len(grid[0]) - 1]


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.minPathSum(
            [[1, 3, 1], [1, 5, 1], [4, 2, 1]]), 7)
