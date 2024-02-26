from typing import List


class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        row = [1]
        for i in range(1, rowIndex + 1):
            next_row = [0] * (i + 1)
            for j in range(i + 1):
                val = 0
                if j > 0:
                    val = row[j - 1]
                if j < i:
                    val += row[j]
                next_row[j] = val
            row = next_row
        return row
