struct Solution;

impl Solution {
    pub fn longest_palindrome(s: String) -> String {
        let mut left = 0;
        let mut right = 0;
        let bytes = s.as_bytes();
        let mut dp = vec![vec![false; s.len()]; s.len()];
        for i in 1..=s.len() {
            for j in 0..=s.len() - i {
                match i {
                    1 => {
                        dp[j][j + i - 1] = true;
                    }
                    2 => {
                        dp[j][j + i - 1] = bytes[j] == bytes[j + 1];
                    }
                    _ => {
                        dp[j][j + i - 1] = bytes[j] == bytes[j + i - 1] && dp[j + 1][j + i - 2];
                    }
                }

                if dp[j][j + i - 1] && i > right - left + 1 {
                    left = j;
                    right = j + i - 1;
                }
            }
        }

        String::from_utf8_lossy(&bytes[left..right + 1]).to_string()
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::longest_palindrome(String::from("babad")),
        String::from("bab")
    );
    assert_eq!(
        Solution::longest_palindrome(String::from("cbbd")),
        String::from("bb")
    );
}
