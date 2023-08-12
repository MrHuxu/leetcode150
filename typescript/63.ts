function uniquePathsWithObstacles(obstacleGrid: number[][]): number {
    const dp: number[][] = new Array(obstacleGrid.length).fill([]).map(() => new Array(obstacleGrid[0].length).fill(0));
    for (let i = 0; i < obstacleGrid.length; i++) {
        for (let j = 0; j < obstacleGrid[0].length; j++) {
            if (obstacleGrid[i][j] === 1) {
                dp[i][j] = 0;
                continue;
            }
            if (i === 0 && j === 0) {
                dp[i][j] = 1;
                continue;
            }
            if (i === 0) {
                dp[i][j] = dp[i][j - 1];
                continue;
            }
            if (j === 0) {
                dp[i][j] = dp[i - 1][j];
                continue;
            }
            dp[i][j] = dp[i - 1][j] + dp[i][j - 1];
        }
    }
    return dp[obstacleGrid.length - 1][obstacleGrid[0].length - 1];
};

test('63', () => {
    expect(uniquePathsWithObstacles([[0, 0, 0], [0, 1, 0], [0, 0, 0]])).toBe(2);
})