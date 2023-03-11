function minDistance(word1: string, word2: string): number {
    const dp: number[][] = Array(word1.length + 1);
    for (let i = 0; i < dp.length; i++) dp[i] = Array(word2.length + 1);

    for (let i = 0; i <= word1.length; i++) {
        for (let j = 0; j <= word2.length; j++) {
            if (i === 0 && j === 0) {
                dp[i][j] = 0;
                continue;
            }
            if (i === 0 || j === 0) {
                dp[i][j] = i + j;
                continue;
            }

            dp[i][j] = Math.min(
                dp[i][j - 1] + 1, dp[i - 1][j] + 1,
                dp[i - 1][j - 1] + (word1[i - 1] === word2[j - 1] ? 0 : 1)
            );
        }
    }

    return dp[word1.length][word2.length];
};

test('72', () => {
    expect(minDistance('horse', 'ros')).toBe(3);
    expect(minDistance('intention', 'execution')).toBe(5);
});