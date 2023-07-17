import unittest
from math import ceil


class Solution:
    def mySqrt(self, x: int) -> int:
        if x == 0 or x == 1:
            return x

        left, right = 1, ceil(x / 2)
        while left <= right:
            mid = ceil((left + right) / 2)
            pow = mid * mid
            if pow == x:
                return mid
            elif pow > x:
                right = mid - 1
            else:
                left = mid + 1

        return right


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.mySqrt(4), 2)
        self.assertEqual(solution.mySqrt(8), 2)
        self.assertEqual(solution.mySqrt(9), 3)
