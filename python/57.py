import unittest
from typing import List


class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        intervals.append(newInterval)
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
        self.assertEqual(solution.insert(
            [[1, 3], [6, 9]], [2, 5]), [[1, 5], [6, 9]])
