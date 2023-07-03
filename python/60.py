import unittest


class Solution:
    COUNT = [0, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880]

    def getPermutation(self, n: int, k: int) -> str:
        ret = ''
        used = [False] * (n + 1)
        for i in range(n):
            for j in range(1, n + 1):
                if not used[j]:
                    if k == 1:
                        ret += str(j)
                        used[j] = True
                        break
                    else:
                        if k <= self.COUNT[n - i - 1]:
                            ret += str(j)
                            used[j] = True
                            break
                        else:
                            k -= self.COUNT[n - i - 1]
        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual('213', solution.getPermutation(3, 3))
        self.assertEqual('2314', solution.getPermutation(4, 9))
