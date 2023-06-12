import unittest
from typing import Dict

map_sym_to_num: Dict[str, int] = {
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}


class Solution:
    def romanToInt(self, s: str) -> int:
        ret = 0

        for i in range(len(s)):
            if s[i] == 'V' or s[i] == 'X':
                j = i - 1
                while j >= 0 and s[j] == 'I':
                    ret -= map_sym_to_num['I'] * 2
                    j -= 1
            elif s[i] == 'L' or s[i] == 'C':
                j = i - 1
                while j >= 0 and s[j] == 'X':
                    ret -= map_sym_to_num['X'] * 2
                    j -= 1
            elif s[i] == 'D' or s[i] == 'M':
                j = i - 1
                while j >= 0 and s[j] == 'C':
                    ret -= map_sym_to_num['C'] * 2
                    j -= 1
            ret += map_sym_to_num[s[i]]

        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.romanToInt('III'), 3)
        self.assertEqual(solution.romanToInt('IV'), 4)
        self.assertEqual(solution.romanToInt('IX'), 9)
