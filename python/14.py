import unittest
from typing import List


class Solution:
    def longestCommonPrefix(self, strs: List[str]) -> str:
        if not strs:
            return ''
        prefix = strs[0]
        for i in range(1, len(strs)):
            while strs[i].find(prefix) != 0:
                prefix = prefix[:-1]
        return prefix


class TestSolution(unittest.TestCase):
    def test_longestCommonPrefix(self):
        s = Solution()
        self.assertEqual(s.longestCommonPrefix(
            ['flower', 'flow', 'flight']), 'fl')
        self.assertEqual(s.longestCommonPrefix(['dog', 'racecar', 'car']), '')
