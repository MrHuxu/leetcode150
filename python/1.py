import unittest
from typing import Dict, List


class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        mapNumToPos: Dict[int, int] = {}
        for idx in range(len(nums)):
            if target - nums[idx] in mapNumToPos:
                return [mapNumToPos[target - nums[idx]], idx]
            mapNumToPos[nums[idx]] = idx
        return []


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertListEqual(solution.twoSum([2, 7, 11, 15], 9), [0, 1])
        self.assertListEqual(solution.twoSum([3, 2, 4], 6), [1, 2])
        self.assertListEqual(solution.twoSum([3, 3], 6), [0, 1])
