struct Solution;

impl Solution {
    pub fn my_sqrt(x: i32) -> i32 {
        let x = x as i64;

        let mut left = 0;
        let mut right = x;
        while left <= right {
            let mid = (left + right) / 2;

            if mid * mid == x || (mid * mid < x && (mid + 1) * (mid + 1) > x) {
                return mid as i32;
            }

            if mid * mid < x {
                left = mid + 1;
            } else {
                right = mid - 1;
            }
        }

        1
    }
}

#[test]
fn test() {
    assert_eq!(Solution::my_sqrt(0), 0);
    assert_eq!(Solution::my_sqrt(1), 1);
    assert_eq!(Solution::my_sqrt(4), 2);
    assert_eq!(Solution::my_sqrt(8), 2);
    assert_eq!(Solution::my_sqrt(2147395599), 46339);
}
