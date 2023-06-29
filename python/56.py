import unittest
from typing import List


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        intervals.sort(key=lambda x: x[0])

        ret: List[List[int]] = []

        for interval in intervals:
            if not ret or ret[-1][1] < interval[0]:
                ret.append(interval)
            else:
                ret[-1][1] = max(ret[-1][1], interval[1])

        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertListEqual(solution.merge([[1, 3], [2, 6], [8, 10], [15, 18]]), [
            [1, 6], [8, 10], [15, 18]])
        self.assertListEqual(solution.merge([[1, 4], [0, 4]]), [[0, 4]])
