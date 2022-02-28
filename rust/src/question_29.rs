struct Solution;

use std::i32::{MAX, MIN};

impl Solution {
    pub fn divide(dividend: i32, divisor: i32) -> i32 {
        if dividend == 0 {
            return 0;
        }

        let sym = if (dividend as i64) * (divisor as i64) > 0 {
            1
        } else {
            -1
        };
        let mut ret = 0;

        let mut dividend: i64 = (dividend as i64).abs();
        let divisor: i64 = (divisor as i64).abs();
        while dividend >= divisor {
            let mut multi = divisor;
            while dividend >= multi {
                multi += multi;
            }

            dividend -= multi / 2;
            ret += multi / 2 / divisor;
        }

        if ret * sym > MAX as i64 {
            MAX
        } else if ret & sym < MIN as i64 {
            MIN
        } else {
            (ret * sym) as i32
        }
    }
}

#[test]
fn test() {
    assert_eq!(Solution::divide(10, 3), 3);
    assert_eq!(Solution::divide(7, -3), -2);
    assert_eq!(Solution::divide(-2147483648, -1), 2147483647);
}
