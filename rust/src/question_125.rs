struct Solution;

impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        let mut head = 0;
        let mut tail = s.len() - 1;

        let bytes = s.as_bytes();
        while head < tail {
            match (Self::conv(bytes[head]), Self::conv(bytes[tail])) {
                (None, _) => head += 1,
                (_, None) => tail -= 1,
                (Some(val1), Some(val2)) => {
                    if val1 != val2 {
                        return false;
                    }

                    head += 1;
                    tail -= 1;
                }
            }
        }

        true
    }

    fn conv(byte: u8) -> Option<u8> {
        match byte as char {
            'a'..='z' | '0'..='9' => Some(byte),
            'A'..='Z' => Some(byte - 'A' as u8 + 'a' as u8),
            _ => None,
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::is_palindrome(String::from("A man, a plan, a canal: Panama")),
        true
    );
    assert_eq!(Solution::is_palindrome(String::from("race a car")), false);
    assert_eq!(Solution::is_palindrome(String::from(" ")), true);
    assert_eq!(Solution::is_palindrome(String::from("0P")), false);
}
