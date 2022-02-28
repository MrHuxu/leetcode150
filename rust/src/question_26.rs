struct Solution;

impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        let mut ret = 0;

        let mut curr = -200;
        for i in 0..nums.len() {
            if nums[i] != curr {
                curr = nums[i];
                nums[ret as usize] = curr;
                ret += 1;
            }
        }

        ret
    }
}

#[test]
fn test() {
    let nums = &mut vec![1, 1, 2];
    assert_eq!(Solution::remove_duplicates(nums), 2);
    assert_eq!(nums[..2], vec![1, 2]);

    let nums = &mut vec![0, 0, 1, 1, 1, 2, 2, 2, 3, 4];
    assert_eq!(Solution::remove_duplicates(nums), 5);
    assert_eq!(nums[..5], vec![0, 1, 2, 3, 4]);
}
