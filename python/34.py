import unittest
from typing import List


class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        ret = [-1, -1]

        if not nums:
            return ret

        left, right = 0, len(nums) - 1
        while left < right:
            mid = (left + right) // 2

            if nums[mid] < target:
                left = mid + 1
            else:
                right = mid
        if nums[(left + right) // 2] == target:
            ret[0] = (left + right) // 2

        left, right = 0, len(nums) - 1
        while left <= right:
            mid = (left + right) // 2

            if nums[mid] > target:
                right = mid - 1
            else:
                left = mid + 1
        if nums[(left + right) // 2] == target:
            ret[1] = (left + right) // 2

        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.searchRange([5, 7, 7, 8, 8, 10], 8), [3, 4])
        self.assertEqual(solution.searchRange(
            [5, 7, 7, 8, 8, 10], 6), [-1, -1])
        self.assertEqual(solution.searchRange([], 0), [-1, -1])
