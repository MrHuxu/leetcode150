struct Solution;

impl Solution {
    pub fn partition(s: String) -> Vec<Vec<String>> {
        let bytes = s.as_bytes();
        let mut dp = vec![vec![false; s.len()]; s.len()];

        for l in 1..=s.len() {
            for i in 0..=s.len() - l {
                let j = i + l - 1;

                dp[i][j] = match l {
                    1 => true,
                    2 => bytes[i] == bytes[j],
                    _ => dp[i + 1][j - 1] && bytes[i] == bytes[j],
                };
            }
        }

        let mut ret = Vec::new();
        Self::helper(&dp, &bytes.to_vec(), 0, vec![], &mut ret);
        ret
    }

    fn helper(
        dp: &Vec<Vec<bool>>,
        bytes: &Vec<u8>,
        idx: usize,
        arr: Vec<String>,
        ret: &mut Vec<Vec<String>>,
    ) {
        if idx == bytes.len() {
            ret.push(arr);
            return;
        }

        for i in idx..bytes.len() {
            if dp[idx][i] {
                let mut tmp = arr.clone();
                tmp.push(String::from_utf8_lossy(&bytes[idx..i + 1]).to_string());
                Self::helper(dp, bytes, i + 1, tmp, ret);
            }
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::partition(String::from("aab")),
        vec![vec!["a", "a", "b"], vec!["aa", "b"]]
    );
}
