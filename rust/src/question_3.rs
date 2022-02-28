struct Solution;

use std::cmp::max;
use std::collections::HashMap;

impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        if s.len() <= 1 {
            return s.len() as i32;
        }

        let mut ret = 1;

        let bytes = s.as_bytes();
        let mut existence: HashMap<u8, i32> = HashMap::new();
        existence.insert(bytes[0], 0);
        let mut left = 0;
        for (right, ch) in bytes.iter().enumerate() {
            if right == 0 {
                continue;
            }

            match existence.get(ch) {
                None => {
                    existence.insert(*ch, 0);
                }
                Some(_val) => {
                    while bytes[left] != *ch {
                        existence.remove(&bytes[left]);
                        left += 1;
                    }
                    left += 1;
                }
            }

            ret = max(ret, (right - left + 1) as i32);
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::length_of_longest_substring(String::from("abcabcbb")),
        3
    );
    assert_eq!(
        Solution::length_of_longest_substring(String::from("bbbbb")),
        1
    );
    assert_eq!(
        Solution::length_of_longest_substring(String::from("pwwkew")),
        3
    );
}
