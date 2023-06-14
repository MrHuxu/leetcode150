import unittest
from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] < nums[0]:
                right = mid - 1
            else:
                left = mid + 1

        rotate_idx = (left + right) // 2
        if target >= nums[0]:
            left, right = 0, rotate_idx
        else:
            left, right = rotate_idx + 1, len(nums) - 1

        while left <= right:
            mid = (left + right) // 2

            if nums[mid] == target:
                return mid
            elif nums[mid] < target:
                left = mid + 1
            else:
                right = mid - 1

        return -1


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.search([4, 5, 6, 7, 0, 1, 2], 0), 4)
