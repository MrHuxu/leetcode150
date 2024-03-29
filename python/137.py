import unittest
from typing import List


class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        ret = 0
        for i in range(32):
            cnt = 0
            for num in nums:
                cnt += (num >> i) & 1
            if cnt % 3 != 0:
                ret |= 1 << i
        if ret >= 2**31:
            ret -= 2**32
        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.singleNumber([2, 2, 3, 2]), 3)
        self.assertEqual(solution.singleNumber([0, 1, 0, 1, 0, 1, 99]), 99)
        self.assertEqual(solution.singleNumber(
            [30000, 500, 100, 30000, 100, 30000, 100]), 500)
        self.assertEqual(solution.singleNumber(
            [-2, -2, 1, 1, 4, 1, 4, 4, -4, -2]), -4)
