import unittest
from typing import Dict


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        ret = 0

        front = 0
        cnt: Dict[str, int] = dict()
        for i in range(len(s)):
            ch = s[i]
            if ch not in cnt:
                ret = max(ret, i - front + 1)
            else:
                while ch in cnt:
                    del cnt[s[front]]
                    front += 1
            cnt[ch] = 1
        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.lengthOfLongestSubstring("abcabcbb"), 3)
        self.assertEqual(solution.lengthOfLongestSubstring("tmmzuxt"), 5)
