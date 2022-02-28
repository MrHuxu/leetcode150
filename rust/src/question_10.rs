struct Solution;

impl Solution {
    pub fn is_match(s: String, p: String) -> bool {
        let bytes_s = s.as_bytes();
        let bytes_p = p.as_bytes();
        let mut dp = vec![vec![false; s.len() + 1]; p.len() + 1];
        for i in 0..=p.len() {
            for j in 0..=s.len() {
                dp[i][j] = match (i, j) {
                    (0, 0) => true,
                    (0, _) => false,
                    (_, 0) => bytes_p[i - 1] == '*' as u8 && dp[i - 2][0],
                    (_, _) => match bytes_p[i - 1] as char {
                        '*' => {
                            dp[i - 2][j]
                                || ((bytes_p[i - 2] == bytes_s[j - 1]
                                    || bytes_p[i - 2] == '.' as u8)
                                    && dp[i][j - 1])
                        }
                        '.' => dp[i - 1][j - 1],
                        _ => bytes_p[i - 1] == bytes_s[j - 1] && dp[i - 1][j - 1],
                    },
                };
            }
        }
        dp[p.len()][s.len()]
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::is_match(String::from("aa"), String::from("a")),
        false
    );

    assert_eq!(
        Solution::is_match(String::from("aa"), String::from("a*")),
        true
    );

    assert_eq!(
        Solution::is_match(String::from("ab"), String::from(".*")),
        true
    );

    assert_eq!(
        Solution::is_match(String::from("mississippi"), String::from("mis*is*p*.")),
        false
    );
}
