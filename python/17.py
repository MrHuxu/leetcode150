import unittest
from typing import Dict, List

MAP_DIGITS_TO_LETTERS : Dict[str, str] = {
    '2': 'abc',
    '3': 'def',
    '4': 'ghi',
    '5': 'jkl',
    '6': 'mno',
    '7': 'pqrs',
    '8': 'tuv',
    '9': 'wxyz'
}


class Solution:
    def letterCombinations(self, digits: str) -> List[str]:
        return self.helper(0, digits) if digits else []
    
    def helper(self, idx: int, digits: str) -> List[str]:
        if idx == len(digits):
            return ['']
        
        ret: List[str] = []
        for next in self.helper(idx + 1, digits):
            for letter in MAP_DIGITS_TO_LETTERS[digits[idx]]:
                ret.append(letter + next)
        return ret


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual([], solution.letterCombinations(''))
        self.assertEqual(['ad', 'bd', 'cd', 'ae', 'be', 'ce', 'af', 'bf', 'cf'], solution.letterCombinations('23'))
        self.assertEqual(['a', 'b', 'c'], solution.letterCombinations('2'))