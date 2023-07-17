import unittest
from typing import List


class Solution:
    def simplifyPath(self, path: str) -> str:
        stack: List[str] = []
        for p in path.split('/'):
            if not p or p == '.':
                continue
            elif p == '..':
                stack = stack[:-1]
            else:
                stack.append(p)
        return '/' + '/'.join(stack)


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual('/home', solution.simplifyPath('/home/'))
