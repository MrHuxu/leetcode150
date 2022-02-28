struct Solution;

impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        if x < 0 {
            return false;
        }

        let mut mut_x = x;
        let mut reversed = 0;
        while mut_x != 0 {
            reversed = reversed * 10 + mut_x % 10;
            mut_x /= 10;
        }

        reversed == x
    }
}

#[test]
fn test() {
    assert_eq!(Solution::is_palindrome(123), false);
    assert_eq!(Solution::is_palindrome(121), true);
    assert_eq!(Solution::is_palindrome(-121), false);
}
