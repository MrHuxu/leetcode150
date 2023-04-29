function longestPalindrome(s: string): string {
    if (!s.length) return s;

    const dp: boolean[][] = new Array(s.length).fill(false).map(
        () => new Array(s.length).fill(false)
    );

    let resultL = 0;
    let resultR = 0;
    for (let l = 1; l <= s.length; l++) {
        for (let i = 0; i + l <= s.length; i++) {
            const j = i + l - 1;
            dp[i][j] = s[i] === s[j] && (l <= 2 || dp[i + 1][j - 1]);
            if (dp[i][j]) {
                resultL = i;
                resultR = j;
            }
        }
    }

    return s.slice(resultL, resultR + 1);
};

test('5', () => {
    expect(longestPalindrome('babad')).toBe('aba');
    expect(longestPalindrome('cbbd')).toBe('bb');
});