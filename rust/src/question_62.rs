struct Solution;

impl Solution {
    pub fn unique_paths(m: i32, n: i32) -> i32 {
        let m = m as usize;
        let n = n as usize;

        let mut dp = vec![vec![0; n]; m];
        dp[0][0] = 1;
        for i in 0..m {
            for j in 0..n {
                if i == 0 && j == 0 {
                    continue;
                }

                dp[i][j] =
                    if i > 0 { dp[i - 1][j] } else { 0 } + if j > 0 { dp[i][j - 1] } else { 0 }
            }
        }

        dp[m - 1][n - 1]
    }
}

#[test]
fn test() {
    assert_eq!(Solution::unique_paths(3, 7), 28);
    assert_eq!(Solution::unique_paths(3, 2), 3);
}
