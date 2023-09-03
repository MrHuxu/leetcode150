function uniquePaths(m: number, n: number): number {
    const dp = new Array(m).fill(0).map(() => new Array(n).fill(0));
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (i == 0 && j == 0)
                dp[i][j] = 1;
            if (i >= 1 && i < m)
                dp[i][j] += dp[i - 1][j]
            if (j >= 1)
                dp[i][j] += dp[i][j - 1]
        }
    }
    return dp[m - 1][n - 1];
};

test('62', () => {
    expect(uniquePaths(3, 2)).toBe(3);
    expect(uniquePaths(7, 3)).toBe(28);
});