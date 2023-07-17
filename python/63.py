import unittest
from typing import List


class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        dp: List[List[int]] = [[0] * len(obstacleGrid[0])
                               for _ in range(len(obstacleGrid))]

        for i in range(len(obstacleGrid)):
            for j in range(len(obstacleGrid[0])):
                if obstacleGrid[i][j] == 1:
                    continue

                if i == 0 and j == 0:
                    dp[i][j] = 1

                if i - 1 >= 0 and i - 1 < len(obstacleGrid) and obstacleGrid[i - 1][j] != 1:
                    dp[i][j] += dp[i - 1][j]

                if j - 1 >= 0 and obstacleGrid[i][j - 1] != 1:
                    dp[i][j] += dp[i][j - 1]

        return dp[len(obstacleGrid) - 1][len(obstacleGrid[0]) - 1]


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.uniquePathsWithObstacles(
            [[0, 0, 0], [0, 1, 0], [0, 0, 0]]), 2)
