struct Solution;

impl Solution {
    pub fn climb_stairs(n: i32) -> i32 {
        let n = n as usize;
        let mut dp = vec![0; n + 1];
        dp[0] = 1;
        for i in 0..=n {
            if i + 1 <= n {
                dp[i + 1] += dp[i];
            }
            if i + 2 <= n {
                dp[i + 2] += dp[i];
            }
        }
        dp[n]
    }
}

#[test]
fn test() {
    assert_eq!(Solution::climb_stairs(2), 2);
    assert_eq!(Solution::climb_stairs(3), 3);
}
