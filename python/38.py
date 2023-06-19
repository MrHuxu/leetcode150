import unittest


class Solution:
    def countAndSay(self, n: int) -> str:
        ret = '1'
        for _ in range(n - 1):
            next = ''
            ch = ret[0]
            count = 1
            for i in range(1, len(ret)):
                if ret[i] == ch:
                    count += 1
                    continue

                next = f'{next}{count}{ch}'
                ch = ret[i]
                count = 1
            ret = f'{next}{count}{ch}'
        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.countAndSay(4), '1211')
