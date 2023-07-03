import unittest


class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        dp = [[0 for _ in range(n)] for _ in range(m)]
        for i in range(m):
            for j in range(n):
                if i == 0 and j == 0:
                    dp[i][j] = 1
                if i >= 1 and i < m:
                    dp[i][j] += dp[i - 1][j]
                if j >= 1:
                    dp[i][j] += dp[i][j - 1]
        return dp[m - 1][n - 1]


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.uniquePaths(3, 2), 3)
        self.assertEqual(solution.uniquePaths(7, 3), 28)
