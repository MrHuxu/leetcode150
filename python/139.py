import unittest
from typing import List, Set


class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        if not wordDict:
            return s == ''

        min_len, max_len = len(wordDict[0]), len(wordDict[0])
        for word in wordDict:
            min_len = min(min_len, len(word))
            max_len = max(max_len, len(word))

        exist: Set[str] = set(wordDict)
        dp: List[bool] = [False] * (len(s) + 1)
        dp[0] = True
        for i in range(1, len(s) + 1):
            if i < min_len:
                dp[i] = False
                continue

            found = False
            for j in range(min_len, min(i, max_len) + 1):
                if dp[i - j] and s[i - j:i] in exist:
                    found = True
                    break
            dp[i] = found

        return dp[-1]


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(solution.wordBreak('leetcode', ['leet', 'code']))
        self.assertTrue(solution.wordBreak('applepenapple', ['apple', 'pen']))
        self.assertFalse(solution.wordBreak(
            'catsandog', ['cats', 'dog', 'sand', 'and', 'cat']))
