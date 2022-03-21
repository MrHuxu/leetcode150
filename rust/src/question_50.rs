struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn my_pow(x: f64, n: i32) -> f64 {
        if n == 0 || x == 1.0 {
            return 1.0;
        }

        let sym = n > 0;
        let mut abs_n = if sym { n as i64 } else { 0 - n as i64 };

        let mut pow_2_val = HashMap::new();
        let mut pow_2 = 1;
        let mut val = x;
        while pow_2 <= abs_n {
            pow_2_val.insert(pow_2, val);
            pow_2 *= 2;
            val *= val;
        }

        let mut ret = 1.0;
        while abs_n != 0 {
            if pow_2 > abs_n {
                pow_2 /= 2;
                continue;
            }

            abs_n -= pow_2;
            ret *= *pow_2_val.get(&pow_2).unwrap();
            pow_2 /= 2;
        }

        if !sym {
            return 1.0 / ret;
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(Solution::my_pow(2.0, 10), 1024.0);
    assert_eq!(Solution::my_pow(2.1, 3), 9.261000000000001);
    assert_eq!(Solution::my_pow(2.0, -2), 0.25);
    assert_eq!(Solution::my_pow(0.00001, 2147483647), 0.0);
    assert_eq!(Solution::my_pow(2.0, -2147483648), 0.0);
}
