import unittest
from typing import List, Set


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        return self.dfs(set(), nums)

    def dfs(self, used: Set[int], nums: List[int]) -> List[List[int]]:
        ret: List[List[int]] = []
        for num in nums:
            if num in used:
                continue
            used.add(num)
            for tmp in self.dfs(used, nums):
                ret.append([num] + tmp)
            used.remove(num)
        return ret if ret else [[]]


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.permute([1, 2, 3]), [[1, 2, 3], [1, 3, 2], [
                         2, 1, 3], [2, 3, 1], [3, 1, 2], [3, 2, 1]])
