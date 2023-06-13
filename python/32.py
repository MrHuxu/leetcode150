import unittest
from typing import List


class Solution:
    def longestValidParentheses(self, s: str) -> int:
        ret = 0

        dp: List[int] = [0 for _ in range(len(s))]
        for i in range(1, len(s)):
            if s[i] == ')':
                if i - 1 - dp[i - 1] >= 0 and s[i - 1 - dp[i - 1]] == '(':
                    dp[i] = dp[i - 1] + 2
                    if i - dp[i] >= 0:
                        dp[i] += dp[i - dp[i]]
                    ret = max(ret, dp[i])

        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.longestValidParentheses("(()"), 2)
