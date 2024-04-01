class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        return len(s.split()[-1])
        return next((len(word) for word in reversed(s.split(" ")) if word), 0)
