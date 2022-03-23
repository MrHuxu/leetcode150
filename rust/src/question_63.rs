struct Solution;

impl Solution {
    pub fn unique_paths_with_obstacles(obstacle_grid: Vec<Vec<i32>>) -> i32 {
        let mut dp = vec![vec![0; obstacle_grid[0].len()]; obstacle_grid.len()];

        dp[0][0] = if obstacle_grid[0][0] == 1 { 0 } else { 1 };
        for i in 0..obstacle_grid.len() {
            for j in 0..obstacle_grid[0].len() {
                if (i == 0 && j == 0) || obstacle_grid[i][j] == 1 {
                    continue;
                }

                dp[i][j] =
                    if i > 0 { dp[i - 1][j] } else { 0 } + if j > 0 { dp[i][j - 1] } else { 0 }
            }
        }

        dp[obstacle_grid.len() - 1][obstacle_grid[0].len() - 1]
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::unique_paths_with_obstacles(vec![vec![0, 0, 0], vec![0, 1, 0], vec![0, 0, 0]]),
        2
    );
    assert_eq!(
        Solution::unique_paths_with_obstacles(vec![vec![0, 1], vec![0, 0]]),
        1
    );
}
