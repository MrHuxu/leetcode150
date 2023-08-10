import unittest
from typing import List, Optional


class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        if target == nums[0]:
            return True

        rotate_idx = self.find_rotate_idx(nums, 0, len(nums) - 1)
        if rotate_idx is None:
            return False

        left, right = (0, rotate_idx) if target > nums[0] else (
            rotate_idx + 1, len(nums) - 1)
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] == target:
                return True
            elif target < nums[mid]:
                right = mid - 1
            else:
                left = mid + 1
        return False

    def find_rotate_idx(self, nums: List[int], left: int, right: int) -> Optional[int]:
        while left <= right:
            mid = (left + right) // 2
            if mid == len(nums) - 1 or nums[mid] > nums[mid + 1]:
                return mid
            if nums[mid] < nums[0]:
                right = mid - 1
                continue
            if nums[mid] > nums[0]:
                left = mid + 1
                continue
            idx = self.find_rotate_idx(nums, left, mid - 1)
            if idx is not None:
                return idx
            return self.find_rotate_idx(nums, mid + 1, right)
        return None


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        # self.assertTrue(solution.search([2, 5, 6, 0, 0, 1, 2], 0))
        # self.assertFalse(solution.search([2, 5, 6, 0, 0, 1, 2], 3))
        self.assertTrue(solution.search([1, 0, 1, 1, 1], 0))
