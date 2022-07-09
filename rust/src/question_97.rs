pub struct Solution;

impl Solution {
    pub fn is_interleave(s1: String, s2: String, s3: String) -> bool {
        if s1.len() + s2.len() != s3.len() {
            return false;
        }

        let mut dp = vec![vec![false; s2.len() + 1]; s1.len() + 1];

        let bytes1 = s1.as_bytes();
        let bytes2 = s2.as_bytes();
        let bytes3 = s3.as_bytes();

        for i in 0..=s1.len() {
            for j in 0..=s2.len() {
                dp[i][j] = match (i, j) {
                    (0, 0) => true,
                    (0, _) => s2[..j] == s3[..j],
                    (_, 0) => s1[..i] == s3[..i],
                    (_, _) => {
                        (bytes3[i + j - 1] == bytes1[i - 1] && dp[i - 1][j])
                            || (bytes3[i + j - 1] == bytes2[j - 1] && dp[i][j - 1])
                    }
                }
            }
        }

        dp[s1.len()][s2.len()]
    }
}

#[test]
fn test() {
    assert_eq!(
        true,
        Solution::is_interleave(
            String::from("aabcc"),
            String::from("dbbca"),
            String::from("aadbbcbcac")
        )
    );
    assert_eq!(
        false,
        Solution::is_interleave(
            String::from("aabcc"),
            String::from("dbbca"),
            String::from("aadbbbaccc")
        )
    );
    assert_eq!(
        true,
        Solution::is_interleave(String::from(""), String::from(""), String::from(""))
    );
}
