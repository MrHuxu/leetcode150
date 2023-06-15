import unittest
from typing import Dict


class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s

        m: Dict[int, str] = {}
        vertical = True
        line = 0
        for char in s:
            m[line] = m.get(line, '') + char
            line += 1 if vertical else -1

            if line == numRows - 1:
                vertical = False
            elif line == 0:
                vertical = True

        ret = ''
        for i in range(numRows):
            ret += m.get(i, '')
        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.convert(
            'PAYPALISHIRING', 3), 'PAHNAPLSIIGYIR')
