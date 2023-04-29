import unittest
from typing import Dict, List


class Solution:
    mapRightToLeft: Dict[str, str] = dict(zip(')]}', '([{'))

    def isValid(self, s: str) -> bool:
        stack: List[str] = []
        for char in s:
            if not stack:
                stack.append(char)
                continue

            if self.mapRightToLeft.get(char) == stack[-1]:
                stack.pop()
            else:
                stack.append(char)
        return not stack


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(solution.isValid('()'))
        self.assertFalse(solution.isValid('(]'))
