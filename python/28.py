class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        for i in range(len(haystack) - len(needle) + 1):
            mismatch = False
            for j in range(len(needle)):
                if haystack[i + j] != needle[j]:
                    mismatch = True
                    break
            if not mismatch:
                return i
        return -1
