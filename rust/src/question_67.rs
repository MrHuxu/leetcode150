struct Solution;

use std::cmp::max;

impl Solution {
    pub fn add_binary(a: String, b: String) -> String {
        let mut ret = Vec::new();

        let bytes_a = a.as_bytes();
        let bytes_b = b.as_bytes();
        let mut carry = 0;
        for i in 0..max(a.len(), b.len()) {
            let sum = match (i < a.len(), i < b.len()) {
                (false, false) => break,
                (false, true) => bytes_b[b.len() - 1 - i] - '0' as u8 + carry,
                (true, false) => bytes_a[a.len() - 1 - i] - '0' as u8 + carry,
                (true, true) => {
                    (bytes_a[a.len() - 1 - i] - '0' as u8)
                        + (bytes_b[b.len() - 1 - i] - '0' as u8)
                        + carry
                }
            };

            ret.push(sum % 2 + '0' as u8);
            carry = sum / 2;
        }
        if carry > 0 {
            ret.push(carry + '0' as u8);
        }

        ret.reverse();
        String::from_utf8_lossy(&ret).to_string()
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::add_binary(String::from("11"), String::from("1")),
        String::from("100")
    );
    assert_eq!(
        Solution::add_binary(String::from("1010"), String::from("1011")),
        String::from("10101")
    );
}
