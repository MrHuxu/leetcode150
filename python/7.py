import unittest


class Solution:
    def reverse(self, x: int) -> int:
        return self.helper(0, x) if x >= 0 else -self.helper(0, -x)

    def helper(self, pre: int, num: int) -> int:
        if num == 0:
            return pre
        else:
            ret = self.helper(pre * 10 + num % 10, num // 10)
            return 0 if ret > 2 ** 31 - 1 or ret < -2 ** 31 else ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.reverse(123), 321)
        self.assertEqual(solution.reverse(-123), -321)
        self.assertEqual(solution.reverse(120), 21)
