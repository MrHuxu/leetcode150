import unittest
from typing import List, Set


class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        return self.dfs(nums, set())

    def dfs(self, nums: List[int], used: Set[int]) -> List[List[int]]:
        ret = []

        for i in range(len(nums)):
            if i in used or (i > 0 and nums[i] == nums[i - 1] and i - 1 not in used):
                continue

            used.add(i)
            for next in self.dfs(nums, used):
                ret.append([nums[i]] + next)
            used.remove(i)

        if not ret:
            ret.append([])

        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertListEqual(solution.permuteUnique([1, 1, 2]), [
                             [1, 1, 2], [1, 2, 1], [2, 1, 1]])
