struct Solution;

impl Solution {
    pub fn gray_code(n: i32) -> Vec<i32> {
        let mut ret = vec![0, 1];
        if n == 1 {
            return ret;
        }

        for _i in 2..=n {
            let len = ret.len();
            for j in (0..len).rev() {
                ret.push((ret[j] << 1) + 1);
            }
            for j in 0..len {
                ret[j] = ret[j] << 1;
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(Solution::gray_code(1), vec![0, 1]);
    assert_eq!(Solution::gray_code(2), vec![0, 2, 3, 1]);
    assert_eq!(Solution::gray_code(3), vec![0, 4, 6, 2, 3, 7, 5, 1]);
}
