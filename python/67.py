import unittest
from typing import List


class Solution:
    def addBinary(self, a: str, b: str) -> str:
        l = max(len(a), len(b))
        sum: List[str] = ['0'] * (l + 1)
        carry = 0
        for i in range(l):
            if i < len(a):
                carry += ord(a[len(a) - i - 1]) - ord('0')
            if i < len(b):
                carry += ord(b[len(b) - i - 1]) - ord('0')
            sum[l - i] = str(carry % 2)
            carry //= 2
        if carry == 1:
            sum[0] = '1'
        else:
            sum.pop(0)
        return ''.join(sum)


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual('100', solution.addBinary('11', '1'))
        self.assertEqual('10101', solution.addBinary('1010', '1011'))
