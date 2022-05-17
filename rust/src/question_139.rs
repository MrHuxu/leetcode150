struct Solution;

use std::cmp::{max, min};
use std::collections::HashMap;
impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
        let mut words = HashMap::new();
        let mut max_len = word_dict[0].len();
        let mut min_len = word_dict[0].len();
        for word in word_dict.iter() {
            max_len = max(max_len, word.len());
            min_len = min(min_len, word.len());
            words.insert(word, true);
        }

        let mut dp = vec![false; s.len() + 1];
        dp[0] = true;
        for i in 1..=s.len() {
            if i < min_len {
                dp[i] = false;
                continue;
            }

            let mut found = false;
            for j in min_len..=min(max_len, i) {
                if dp[i - j] && words.contains_key(&s[i - j..i].to_string()) {
                    found = true;
                    break;
                }
            }
            dp[i] = found;
        }

        dp[s.len()]
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::word_break(
            String::from("leetcode"),
            vec![String::from("leet"), String::from("code")]
        ),
        true
    );
    assert_eq!(
        Solution::word_break(
            String::from("applepenapple"),
            vec![String::from("apple"), String::from("pen")]
        ),
        true
    );
    assert_eq!(
        Solution::word_break(
            String::from("catsandog"),
            vec![
                String::from("cats"),
                String::from("dog"),
                String::from("sand"),
                String::from("and"),
                String::from("cat")
            ]
        ),
        false
    );
}
