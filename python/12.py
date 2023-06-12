import unittest
from typing import Dict

map_num_to_sym: Dict[int, Dict[int, str]] = {
    0: {
        1:  'I',
        5:  'V',
        10: 'X',
    },
    1: {
        1:  'X',
        5:  'L',
        10: 'C',
    },
    2: {
        1:  'C',
        5:  'D',
        10: 'M',
    },
    3: {
        1: 'M',
    },
}


class Solution:
    def intToRoman(self, num: int) -> str:
        ret = ''

        level = 0
        while num > 0:
            tmp = num % 10

            if tmp > 0 and tmp <= 3:
                ret = map_num_to_sym[level][1] * tmp + ret
                tmp = 0
            elif tmp == 4:
                ret = map_num_to_sym[level][1] + map_num_to_sym[level][5] + ret
            elif tmp > 4 and tmp <= 8:
                ret = map_num_to_sym[level][1] * (tmp - 5) + ret
                tmp -= 5
                ret = map_num_to_sym[level][5] + ret
            elif tmp == 9:
                ret = map_num_to_sym[level][1] + \
                    map_num_to_sym[level][10] + ret

            level += 1
            num //= 10

        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.intToRoman(3), 'III')
        self.assertEqual(solution.intToRoman(4), 'IV')
        self.assertEqual(solution.intToRoman(9), 'IX')
