import math
import unittest
from typing import Callable, Dict, List


class Solution:
    SYMBOLS: Dict[str, Callable[[int, int], int]] = {
        "+": lambda x, y: x + y,
        "-": lambda x, y: x - y,
        "*": lambda x, y: x * y,
        "/": lambda x, y: math.trunc(x / y),
    }

    def evalRPN(self, tokens: List[str]) -> int:
        stack = []
        for token in tokens:
            if token in self.SYMBOLS:
                y, x = stack.pop(), stack.pop()
                stack.append(self.SYMBOLS[token](x, y))
            else:
                stack.append(int(token))
        return stack[0]


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.evalRPN(["2", "1", "+", "3", "*"]), 9)
        self.assertEqual(solution.evalRPN(["4", "13", "5", "/", "+"]), 6)
        self.assertEqual(
            solution.evalRPN(
                ["10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"]
            ),
            22,
        )
