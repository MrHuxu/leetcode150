class Solution:
    def trap(self, height: List[int]) -> int:
        ret, stack = 0, []
        for i, curr in enumerate(height):
            while stack and height[stack[-1]] < curr:
                bottom = height[stack.pop()]
                if not stack:
                    break
                distance = i - stack[-1] - 1
                top = min(height[stack[-1]], curr)
                ret += distance * (top - bottom)
            stack.append(i)
        return ret
