struct Solution;

impl Solution {
    pub fn num_decodings(s: String) -> i32 {
        let mut dp = vec![0; s.len() + 1];
        dp[0] = 1;

        let bytes = s.as_bytes();
        for i in 1..=s.len() {
            let byte = bytes[i - 1];

            if byte == b'0' {
                if i == 1 || bytes[i - 2] > b'2' || bytes[i - 2] < b'1' {
                    return 0;
                }

                dp[i] = dp[i - 2];
                continue;
            }

            dp[i] = dp[i - 1];

            if i == 1 {
                continue;
            }

            let num = (bytes[i - 2] - b'0') * 10 + byte - b'0';
            if num >= 10 && num <= 26 {
                dp[i] += dp[i - 2];
            }
        }

        dp[s.len()]
    }
}

#[test]
fn test() {
    assert_eq!(Solution::num_decodings(String::from("12")), 2);
    assert_eq!(Solution::num_decodings(String::from("226")), 3);
    assert_eq!(Solution::num_decodings(String::from("06")), 0);
    assert_eq!(Solution::num_decodings(String::from("27")), 1);
}
