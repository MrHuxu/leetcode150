import unittest
from typing import Dict, List


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups: Dict[str, List[str]] = {}
        for s in strs:
            sorted_s = ''.join(sorted(s))
            if sorted_s not in groups:
                groups[sorted_s] = []
            groups[sorted_s].append(s)
        return list(groups.values())


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.groupAnagrams(["eat", "tea", "tan", "ate", "nat", "bat"]), [
                         ['eat', 'tea', 'ate'], ['tan', 'nat'], ['bat']])
