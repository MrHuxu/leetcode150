import unittest
from typing import List


class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        ret = nums[0]

        pre = nums[0]
        for i in range(1, len(nums)):
            pre = nums[i] if pre < 0 else pre + nums[i]
            ret = max(ret, pre)

        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(6, solution.maxSubArray(
            [-2, 1, -3, 4, -1, 2, 1, -5, 4]))
