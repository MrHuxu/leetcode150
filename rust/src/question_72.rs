struct Solution;

use std::cmp::min;

impl Solution {
    pub fn min_distance(word1: String, word2: String) -> i32 {
        let mut dp = vec![vec![0; word2.len() + 1]; word1.len() + 1];

        let bytes1 = word1.as_bytes();
        let bytes2 = word2.as_bytes();
        for i in 0..=word1.len() {
            for j in 0..=word2.len() {
                dp[i][j] = match (i, j) {
                    (0, 0) => 0,
                    (0, _) => j as i32,
                    (_, 0) => i as i32,
                    (_, _) => {
                        let tmp = min(dp[i - 1][j], dp[i][j - 1]) + 1;
                        if bytes1[i - 1] != bytes2[j - 1] {
                            min(tmp, dp[i - 1][j - 1] + 1)
                        } else {
                            min(tmp, dp[i - 1][j - 1])
                        }
                    }
                };
            }
        }

        dp[word1.len()][word2.len()]
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::min_distance(String::from("horse"), String::from("ros")),
        3
    );
    assert_eq!(
        Solution::min_distance(String::from("intention"), String::from("execution")),
        5
    );
}
