import unittest
from typing import List


class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        '''
        Do not return anything, modify board in-place instead.
        '''
        self.dfs(board, 0, 0)

    def dfs(self, board: List[List[str]], i: int, j: int) -> bool:
        if i == 9:
            return True

        if j == 9:
            return self.dfs(board, i + 1, 0)

        if board[i][j] != '.':
            return self.dfs(board, i, j + 1)

        for ch in '123456789':
            if self.validate(board, i, j, ch):
                board[i][j] = ch
                if self.dfs(board, i, j + 1):
                    return True
                board[i][j] = '.'

        return False

    def validate(self, board: List[List[str]], i: int, j: int, ch: str) -> bool:
        for k in range(9):
            if board[i][k] == ch:
                return False
            if board[k][j] == ch:
                return False
            if board[3 * (i // 3) + k // 3][3 * (j // 3) + k % 3] == ch:
                return False
        return True


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        board = [
            ['5', '3', '.', '.', '7', '.', '.', '.', '.'],
            ['6', '.', '.', '1', '9', '5', '.', '.', '.'],
            ['.', '9', '8', '.', '.', '.', '.', '6', '.'],
            ['8', '.', '.', '.', '6', '.', '.', '.', '3'],
            ['4', '.', '.', '8', '.', '3', '.', '.', '1'],
            ['7', '.', '.', '.', '2', '.', '.', '.', '6'],
            ['.', '6', '.', '.', '.', '.', '2', '8', '.'],
            ['.', '.', '.', '4', '1', '9', '.', '.', '5'],
            ['.', '.', '.', '.', '8', '.', '.', '7', '9']
        ]
        solution.solveSudoku(board)
        self.assertListEqual(board, [
            ['5', '3', '4', '6', '7', '8', '9', '1', '2'],
            ['6', '7', '2', '1', '9', '5', '3', '4', '8'],
            ['1', '9', '8', '3', '4', '2', '5', '6', '7'],
            ['8', '5', '9', '7', '6', '1', '4', '2', '3'],
            ['4', '2', '6', '8', '5', '3', '7', '9', '1'],
            ['7', '1', '3', '9', '2', '4', '8', '5', '6'],
            ['9', '6', '1', '5', '3', '7', '2', '8', '4'],
            ['2', '8', '7', '4', '1', '9', '6', '3', '5'],
            ['3', '4', '5', '2', '8', '6', '1', '7', '9']
        ])
