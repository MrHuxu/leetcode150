import unittest
from operator import le
from typing import List


class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        ret: List[int] = []

        level = 0
        m = len(matrix)
        n = len(matrix[0])
        while len(matrix) > 0 and len(ret) < m * n:
            for i in range(level, n - level):
                ret.append(matrix[level][i])
            for i in range(level + 1, m - level - 1):
                ret.append(matrix[i][n - level - 1])
            tmp = n - level - 1
            while level != m - level - 1 and tmp >= level:
                ret.append(matrix[m - level - 1][tmp])
                tmp -= 1
            tmp = m - level - 2
            while level != n - level - 1 and tmp > level:
                ret.append(matrix[tmp][level])
                tmp -= 1
            level += 1

        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertListEqual([1, 2, 3, 6, 9, 8, 7, 4, 5], solution.spiralOrder(
            [[1, 2, 3], [4, 5, 6], [7, 8, 9]]))
        self.assertListEqual([1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7], solution.spiralOrder(
            [[1, 2, 3, 4], [5, 6, 7, 8], [9, 10, 11, 12]]))
