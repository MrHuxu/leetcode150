import unittest
from typing import List


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        if not matrix:
            return False

        left, right = 0, len(matrix) - 1
        row_idx = 0
        while left <= right:
            mid = (left + right) // 2
            if matrix[mid][0] <= target <= matrix[mid][-1]:
                row_idx = mid
                break
            elif target < matrix[mid][0]:
                right = mid - 1
            else:
                left = mid + 1

        left, right = 0, len(matrix[0]) - 1
        while left <= right:
            mid = (left + right) // 2
            if matrix[row_idx][mid] == target:
                return True
            elif matrix[row_idx][mid] < target:
                left = mid + 1
            else:
                right = mid - 1

        return False


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.searchMatrix(
            matrix=[[1, 3, 5, 7], [10, 11, 16, 20], [23, 30, 34, 60]], target=3), True)
