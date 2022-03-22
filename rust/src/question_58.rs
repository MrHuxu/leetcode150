struct Solution;

impl Solution {
    pub fn length_of_last_word(s: String) -> i32 {
        let mut pre_word_len = 0;
        let mut word_len = 0;

        for byte in s.as_bytes() {
            match *byte as char {
                ' ' => {
                    if word_len != 0 {
                        pre_word_len = word_len;
                        word_len = 0;
                    }
                }
                _ => {
                    word_len += 1;
                }
            }
        }

        if word_len == 0 {
            pre_word_len
        } else {
            word_len
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::length_of_last_word(String::from("Hello World")),
        5
    );
    assert_eq!(
        Solution::length_of_last_word(String::from("   fly me   to   the moon  ")),
        4
    );
    assert_eq!(
        Solution::length_of_last_word(String::from("luffy is still joyboy")),
        6
    );
}

/*
Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.
Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.
Example 3:

Input: s = "luffy is still joyboy"
Output: 6 */
