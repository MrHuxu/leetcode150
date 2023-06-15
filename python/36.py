import unittest
from typing import List


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows, cols, rects = [[False] * 10 for _ in range(
            9)], [[False] * 10 for _ in range(9)], [[False] * 10 for _ in range(9)]

        for i in range(9):
            for j in range(9):
                ch = board[i][j]
                if ch == '.':
                    continue

                if rects[i // 3 * 3 + j // 3][ord(ch) - ord('0')]:
                    return False
                rects[i // 3 * 3 + j // 3][ord(ch) - ord('0')] = True

                if rows[i][ord(ch) - ord('0')]:
                    return False
                rows[i][ord(ch) - ord('0')] = True

                if cols[j][ord(ch) - ord('0')]:
                    return False
                cols[j][ord(ch) - ord('0')] = True

        return True


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertTrue(solution.isValidSudoku([
            ['5', '3', '.', '.', '7', '.', '.', '.', '.'],
            ['6', '.', '.', '1', '9', '5', '.', '.', '.'],
            ['.', '9', '8', '.', '.', '.', '.', '6', '.'],
            ['8', '.', '.', '.', '6', '.', '.', '.', '3'],
            ['4', '.', '.', '8', '.', '3', '.', '.', '1'],
            ['7', '.', '.', '.', '2', '.', '.', '.', '6'],
            ['.', '6', '.', '.', '.', '.', '2', '8', '.'],
            ['.', '.', '.', '4', '1', '9', '.', '.', '5'],
            ['.', '.', '.', '.', '8', '.', '.', '7', '9']]))
