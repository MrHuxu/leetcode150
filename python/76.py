from collections import Counter
import unittest


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        counter = Counter(t)
        cnt = len(t)
        result_l, result_r, result_len, found = 0, 0, len(s), False
        left, right = 0, 0
        while right < len(s):
            if s[right] in counter:
                counter[s[right]] -= 1

                if counter[s[right]] >= 0:
                    cnt -= 1

                if cnt == 0:
                    found = True

                while cnt == 0:
                    if right - left + 1 <= result_len:
                        result_l, result_r, result_len = left, right, right - left + 1

                    if s[left] in counter:
                        counter[s[left]] += 1
                        if counter[s[left]] > 0:
                            cnt += 1

                    left += 1
            right += 1
        return s[result_l : result_r + 1] if found else ""


class TestSolution(unittest.TestCase):
    def test(self):
        solution = Solution()
        self.assertEqual(solution.minWindow("ADOBECODEBANC", "ABC"), "BANC")
