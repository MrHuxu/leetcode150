struct Solution;

use std::cmp::min;

impl Solution {
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        let mut dp = vec![vec![0; grid[0].len()]; grid.len()];

        for i in 0..grid.len() {
            for j in 0..grid[0].len() {
                dp[i][j] = grid[i][j]
                    + match (i > 0, j > 0) {
                        (false, false) => 0,
                        (true, false) => dp[i - 1][j],
                        (false, true) => dp[i][j - 1],
                        (true, true) => min(dp[i - 1][j], dp[i][j - 1]),
                    }
            }
        }

        dp[grid.len() - 1][grid[0].len() - 1]
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::min_path_sum(vec![vec![1, 3, 1], vec![1, 5, 1], vec![4, 2, 1]]),
        7
    );
    assert_eq!(
        Solution::min_path_sum(vec![vec![1, 2, 3], vec![4, 5, 6]]),
        12
    );
}
