struct Solution;

impl Solution {
    pub fn num_trees(n: i32) -> i32 {
        let mut dp = vec![0; n as usize + 1];

        for i in 0..=n as usize {
            match i {
                0 => dp[i] = 1,
                _ => {
                    for j in 1..=i {
                        dp[i] += dp[j - 1] * dp[i - j];
                    }
                }
            }
        }

        dp[n as usize]
    }
}

#[test]
fn test() {
    assert_eq!(Solution::num_trees(3), 5);
    assert_eq!(Solution::num_trees(1), 1);
}
