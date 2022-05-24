package question_32;

class Solution {
    public int longestValidParentheses(String s) {
        int ret = 0;

        int[] dp = new int[s.length()];
        char[] chars = s.toCharArray();
        for (int i = 1; i < dp.length; i++) {
            if (chars[i] != ')')
                continue;

            if (i - 1 >= dp[i - 1] && chars[i - 1 - dp[i - 1]] == '(') {
                dp[i] = dp[i - 1] + 2;

                if (i >= dp[i] && dp[i - dp[i]] > 0)
                    dp[i] += dp[i - dp[i]];

                ret = Math.max(ret, dp[i]);
            }
        }

        return ret;
    }
}