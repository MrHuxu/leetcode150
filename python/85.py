class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        if not matrix:
            return 0

        ret = 0
        heights = [0] * len(matrix[0])
        for row in matrix:
            for i, cell in enumerate(row):
                if cell == "0":
                    heights[i] = 0
                else:
                    heights[i] += 1
            ret = max(ret, self.helper(heights))
        return ret

    def helper(self, heights: List[int]) -> int:
        ret = 0
        i, stack = 0, []
        while i < len(heights) or stack:
            if not stack or (i < len(heights) and heights[i] >= heights[stack[-1]]):
                stack.append(i)
                i += 1
                continue

            height = heights[stack.pop()]
            left = -1 if not stack else stack[-1]
            right = i
            area = (right - left - 1) * height
            ret = max(ret, area)
        return ret
