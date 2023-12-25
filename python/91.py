class Solution:
    def numDecodings(self, s: str) -> int:
        if s[0] == '0':
            return 0

        dp = [1] + [0] * (len(s) - 1)
        for i in range(1, len(s)):
            num = int(s[i])
            if 1 <= num <= 9:
                dp[i] += dp[i - 1]
            if 10 <= int(s[i - 1]) * 10 + num <= 26:
                dp[i] += 1 if i == 1 else dp[i - 2]

        return dp[-1]