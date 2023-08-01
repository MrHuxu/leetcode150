import unittest
from typing import List


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        return self.dfs(1, n, k)

    def dfs(self, curr: int, n: int, k: int) -> List[List[int]]:
        if k == 0:
            return [[]]

        ret = []
        for num in range(curr, n + 1):
            for next in self.dfs(num + 1, n, k - 1):
                ret.append([num] + next)
        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertListEqual(solution.combine(
            4, 2), [[1, 2], [1, 3], [1, 4], [2, 3], [2, 4], [3, 4]])
