struct Solution;

use std::i32::{MAX, MIN};

impl Solution {
    pub fn reverse(x: i32) -> i32 {
        let mut ret: i64 = 0;

        let mut mut_x = x as i64;
        while mut_x != 0 {
            ret = ret * 10 + mut_x % 10;
            if ret > MAX as i64 || ret < MIN as i64 {
                return 0;
            }

            mut_x /= 10;
        }

        ret as i32
    }
}

#[test]
fn test() {
    assert_eq!(Solution::reverse(123), 321);
    assert_eq!(Solution::reverse(-123), -321);
    assert_eq!(Solution::reverse(-120), -21);
    assert_eq!(Solution::reverse(120), 21);
    assert_eq!(Solution::reverse(1534236469), 0);
}
