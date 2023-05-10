from typing import List
import unittest


class Solution:
    def generateMatrix(self, n: int) -> List[List[int]]:
        ret: List[List[int]] = [[0 for _ in range(n)] for _ in range(n)]

        num = 0
        for level in range((n + 1) // 2):
            for i in range(level, n - level):
                num += 1
                ret[level][i] = num
            for i in range(level + 1, n - level - 1):
                num += 1
                ret[i][n - level - 1] = num
            tmp = n - level - 1
            while n - level - 1 != level and tmp >= level:
                num += 1
                ret[n - level - 1][tmp] = num
                tmp -= 1
            tmp = n - level - 2
            while n - level - 1 != level and tmp > level:
                num += 1
                ret[tmp][level] = num
                tmp -= 1

        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertListEqual(
            [[1, 2, 3], [8, 9, 4], [7, 6, 5]], solution.generateMatrix(3))
