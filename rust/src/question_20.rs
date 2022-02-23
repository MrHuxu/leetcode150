struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut pairs: HashMap<u8, u8> = HashMap::new();
        pairs.insert(')' as u8, '(' as u8);
        pairs.insert(']' as u8, '[' as u8);
        pairs.insert('}' as u8, '{' as u8);

        let mut stack: Vec<u8> = Vec::new();

        for ch in s.as_bytes().iter() {
            if stack.is_empty() {
                stack.push(*ch);
                continue;
            }

            match *ch as char {
                ')' | ']' | '}' => {
                    if stack.last().unwrap() == pairs.get(ch).unwrap() {
                        stack.pop();
                    } else {
                        return false;
                    }
                }
                _ => {
                    stack.push(*ch);
                }
            }
        }

        stack.len() == 0
    }
}

#[test]
fn test() {
    assert_eq!(Solution::is_valid(String::from("()")), true);
    assert_eq!(Solution::is_valid(String::from("()[]")), true);
    assert_eq!(Solution::is_valid(String::from("(]")), false);
}
