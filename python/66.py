import unittest
from typing import List


class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        carry = 1
        for i in range(len(digits) - 1, -1, -1):
            digits[i] += carry
            carry = digits[i] // 10
            digits[i] %= 10
        if carry == 1:
            digits.insert(0, carry)
        return digits


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.plusOne([1, 2, 3]), [1, 2, 4])
        self.assertEqual(solution.plusOne([4, 3, 2, 1]), [4, 3, 2, 2])
