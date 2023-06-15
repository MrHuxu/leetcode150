import unittest
from typing import List


class Solution:
    def generateParenthesis(self, n: int) -> List[str]:
        return self.helper(0, 0, n, '')

    def helper(self, numLeft: int, matchedLeft: int, n: int, pre: str) -> List[str]:
        if numLeft < n:
            tmp = self.helper(numLeft + 1, matchedLeft, n, pre + '(')
            if matchedLeft < numLeft:
                tmp += self.helper(numLeft, matchedLeft + 1, n, pre + ')')
            return tmp

        if matchedLeft < numLeft:
            return self.helper(numLeft, matchedLeft + 1, n, pre + ')')

        return [pre]


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.generateParenthesis(
            3), ['((()))', '(()())', '(())()', '()(())', '()()()'])
