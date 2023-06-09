import unittest
from typing import List


class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        if n <= 1:
            return

        for i in range((n + 1)//2):
            for j in range(i, n - 1 - i):
                tmp = matrix[i][j]
                matrix[i][j] = matrix[n - 1 - j][i]
                matrix[n - 1 - j][i] = matrix[n - 1 - i][n - 1 - j]
                matrix[n - 1 - i][n - 1 - j] = matrix[j][n - 1 - i]
                matrix[j][n - 1 - i] = tmp


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
        solution.rotate(matrix)
        self.assertListEqual(matrix, [[7, 4, 1], [8, 5, 2], [9, 6, 3]])
