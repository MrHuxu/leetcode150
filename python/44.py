from typing import List
import unittest


class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        dp: List[List[bool]] = [[False] * (len(s) + 1)
                                for _ in range(len(p) + 1)]
        for i in range(len(p) + 1):
            for j in range(len(s) + 1):
                if i == 0 and j == 0:
                    dp[i][j] = True
                elif i == 0:
                    dp[i][j] = False
                elif j == 0:
                    dp[i][j] = p[i - 1] == '*' and dp[i - 1][0]
                else:
                    if p[i - 1] == '*':
                        dp[i][j] = dp[i - 1][j - 1] or dp[i][j - 1] or dp[i - 1][j]
                    elif p[i - 1] == '?':
                        dp[i][j] = dp[i - 1][j - 1]
                    else:
                        dp[i][j] = dp[i - 1][j - 1] and s[j - 1] == p[i - 1]
        return dp[len(p)][len(s)]


class TestSolution(unittest.TestCase):
    def test(self):
        s = Solution()
        self.assertFalse(s.isMatch("aa", "a"))
        self.assertTrue(s.isMatch("aa", "*"))
        self.assertFalse(s.isMatch("cb", "?a"))
        self.assertTrue(s.isMatch("adceb", "*a*b"))
        self.assertFalse(s.isMatch("acdcb", "a*c?b"))
        self.assertTrue(s.isMatch("a", "a*"))
        self.assertFalse(s.isMatch("a", "a*b"))
