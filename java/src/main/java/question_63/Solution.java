package question_63;

class Solution {
    public int uniquePathsWithObstacles(int[][] obstacleGrid) {
        int[][] dp = new int[obstacleGrid.length][obstacleGrid[0].length];

        for (int i = 0; i < dp.length; i++) {
            for (int j = 0; j < dp[i].length; j++) {
                if (obstacleGrid[i][j] == 1) {
                    dp[i][j] = 0;
                    continue;
                }
                if (i == 0 && j == 0) {
                    dp[i][j] = 1;
                    continue;
                }

                dp[i][j] = (i > 0 && obstacleGrid[i - 1][j] == 0 ? dp[i - 1][j] : 0)
                        + (j > 0 && obstacleGrid[i][j - 1] == 0 ? dp[i][j - 1] : 0);
            }
        }

        return dp[dp.length - 1][dp[0].length - 1];
    }
}