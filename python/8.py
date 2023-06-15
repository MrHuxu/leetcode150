import unittest


class Solution:
    def myAtoi(self, s: str) -> int:
        num = 0

        sign = 1
        started = False
        for char in s:
            if started:
                if not char.isdigit():
                    break

                num = (num * 10 + ord(char) - ord('0'))
                if num * sign > 2**31 - 1:
                    return 2**31 - 1
                if num * sign < -2**31:
                    return -2**31
            else:
                if char == ' ':
                    continue
                elif char == '-':
                    sign = -1
                    started = True
                elif char == '+':
                    started = True
                elif char.isdigit():
                    num = ord(char) - ord('0')
                    started = True
                else:
                    return 0

        return num * sign


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.myAtoi('42'), 42)
        self.assertEqual(solution.myAtoi('   -42'), -42)
        self.assertEqual(solution.myAtoi('4193 with words'), 4193)
