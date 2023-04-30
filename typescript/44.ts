function isMatch(s: string, p: string): boolean {
    const dp: boolean[][] = new Array(p.length + 1).fill(false).map(
        () => new Array(s.length + 1).fill(false)
    );

    for (let i = 0; i <= p.length; i++) {
        for (let j = 0; j <= s.length; j++) {
            if (i === 0 && j === 0)
                dp[i][j] = true;
            else if (i === 0)
                dp[i][j] = false;
            else if (j === 0)
                dp[i][j] = p[i - 1] === '*' && dp[i - 1][0];
            else {
                if (p[i - 1] === '*') {
                    dp[i][j] = dp[i - 1][j - 1] || dp[i][j - 1] || dp[i - 1][j];
                } else if (p[i - 1] === '?') {
                    dp[i][j] = dp[i - 1][j - 1];
                } else {
                    dp[i][j] = dp[i - 1][j - 1] && p[i - 1] === s[j - 1];
                }
            }
        }
    }
    return dp[p.length][s.length];
}

test('44', () => {
    expect(isMatch('aa', 'a')).toBe(false);
    expect(isMatch('aa', '*')).toBe(true);
    expect(isMatch('cb', '?a')).toBe(false);
    expect(isMatch('adceb', '*a*b')).toBe(true);
});