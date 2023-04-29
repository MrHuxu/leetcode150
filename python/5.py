import unittest
from typing import List


class Solution:
    def longestPalindrome(self, s: str) -> str:
        if not s:
            return s

        dp: List[List[bool]] = list(
            list(False for _ in range(len(s))) for _ in range(len(s)))

        resultL = 0
        resultR = 0
        for l in range(1, len(s) + 1):
            for i in range(len(s) - l + 1):
                j = i + l - 1
                dp[i][j] = s[i] == s[j] and (l < 3 or dp[i + 1][j - 1])
                if dp[i][j] and l > max:
                    resultL = i
                    resultR = j

        return s[resultL:resultR + 1]


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual("bab", solution.longestPalindrome("babad"))
        self.assertEqual("bb", solution.longestPalindrome("cbbd"))
