import unittest


class Solution:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False

        reversed = 0
        tmp = x
        while tmp > 0:
            reversed = reversed * 10 + tmp % 10
            tmp //= 10

        return reversed == x


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(solution.isPalindrome(121))
        self.assertFalse(solution.isPalindrome(-121))
        self.assertFalse(solution.isPalindrome(10))
