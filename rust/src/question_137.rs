struct Solution;

impl Solution {
    pub fn single_number(nums: Vec<i32>) -> i32 {
        let mut ret = 0;

        for i in 0..32 {
            let mut cnt = 0;

            for num in nums.iter() {
                cnt += (*num >> i) & 1;
            }

            if cnt % 3 != 0 {
                ret = ret | (1 << i);
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(Solution::single_number(vec![1, 1, 1, 2]), 2);
}
