import unittest
from typing import List


class Solution:
    def minDistance(self, word1: str, word2: str) -> int:
        dp: List[List[int]] = [[0] * (len(word2) + 1)
                               for _ in range(len(word1) + 1)]
        for i in range(len(word1) + 1):
            for j in range(len(word2) + 1):
                if i == 0 and j == 0:
                    continue
                elif i == 0:
                    dp[i][j] = j
                elif j == 0:
                    dp[i][j] = i
                else:
                    if word1[i - 1] == word2[j - 1]:
                        dp[i][j] = dp[i - 1][j - 1]
                    else:
                        dp[i][j] = min(dp[i - 1][j - 1], dp[i - 1]
                                       [j], dp[i][j - 1]) + 1

        return dp[-1][-1]


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.minDistance("horse", "ros"), 3)
        self.assertEqual(solution.minDistance("intention", "execution"), 5)
