struct Solution;

use std::cmp;

impl Solution {
    pub fn longest_valid_parentheses(s: String) -> i32 {
        let mut ret = 0;

        let mut dp = vec![0; s.len()];
        let bytes = s.as_bytes();
        for i in 1..s.len() {
            if bytes[i] != ')' as u8 {
                continue;
            }

            if i - 1 >= dp[i - 1] && bytes[i - 1 - dp[i - 1]] == '(' as u8 {
                dp[i] = dp[i - 1] + 2;

                if i >= dp[i] && dp[i - dp[i]] > 0 {
                    dp[i] += dp[i - dp[i]];
                }

                ret = cmp::max(ret, dp[i]);
            }
        }

        ret as i32
    }
}

#[test]
fn test() {
    assert_eq!(Solution::longest_valid_parentheses(String::from("(()")), 2);
    assert_eq!(Solution::longest_valid_parentheses(String::from("(())")), 4);
    assert_eq!(
        Solution::longest_valid_parentheses(String::from(")(()())")),
        6
    );
    assert_eq!(Solution::longest_valid_parentheses(String::from("")), 0);
    assert_eq!(
        Solution::longest_valid_parentheses(String::from("()(())")),
        6
    );
    assert_eq!(
        Solution::longest_valid_parentheses(String::from("()(()")),
        2
    );
}
