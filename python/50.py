import unittest
from typing import Dict


class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n == 0 or x == 1:
            return 1

        sym = n > 0
        n = abs(n)

        pow_2_vals: Dict[int, float] = {}
        pow_2, val = 1, x
        while pow_2 <= n:
            pow_2_vals[pow_2] = val
            pow_2 *= 2
            val *= val

        ret = 1
        while n != 0:
            if pow_2 > n:
                pow_2 //= 2
                continue

            n -= pow_2
            ret *= pow_2_vals[pow_2]
            pow_2 //= 2

        if not sym:
            return 1 / ret
        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.myPow(2, 10), 1024)
        self.assertEqual(solution.myPow(2, -2), 0.25)
        self.assertEqual(solution.myPow(2, -3), 0.125)
