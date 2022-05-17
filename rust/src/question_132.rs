struct Solution;

use std::cmp::min;

impl Solution {
    pub fn minCut(s: String) -> i32 {
        let bytes = s.as_bytes();
        let mut dp1 = vec![vec![false; s.len()]; s.len()];

        for l in 1..=s.len() {
            for i in 0..=s.len() - l {
                let j = i + l - 1;

                dp1[i][j] = match l {
                    1 => true,
                    2 => bytes[i] == bytes[j],
                    _ => dp1[i + 1][j - 1] && bytes[i] == bytes[j],
                };
            }
        }

        let mut dp2 = vec![0; s.len()];
        for i in 1..s.len() {
            dp2[i] = i;

            for j in 0..=i {
                if dp1[j][i] {
                    if j == 0 {
                        dp2[i] = 0;
                        break;
                    } else {
                        dp2[i] = min(dp2[i], dp2[j - 1] + 1);
                    }
                }
            }
        }

        dp2[s.len() - 1] as i32
    }
}

#[test]
fn test() {
    assert_eq!(Solution::minCut(String::from("aab")), 1);
    assert_eq!(Solution::minCut(String::from("a")), 0);
}
