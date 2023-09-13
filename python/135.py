import unittest
from typing import List


class Solution:
    def candy(self, ratings: List[int]) -> int:
        candies = [1] + [0] * (len(ratings) - 1)

        for i in range(1, len(ratings)):
            if ratings[i] > ratings[i - 1]:
                candies[i] = candies[i - 1] + 1
            else:
                candies[i] = 1

        for i in reversed(range(len(ratings) - 1)):
            if ratings[i] > ratings[i + 1]:
                candies[i] = max(candies[i], candies[i + 1] + 1)

        return sum(candies)


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.candy([1, 0, 2]), 5)
