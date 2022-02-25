struct Solution;

impl Solution {
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        let mut ret = 0;

        for i in 0..nums.len() {
            if nums[i] == val {
                continue;
            }

            nums[ret] = nums[i];
            ret += 1;
        }

        ret as i32
    }
}

#[test]
fn test() {
    let nums = &mut vec![3, 2, 2, 3];
    assert_eq!(Solution::remove_element(nums, 3), 2);
    assert_eq!(nums[..2].to_vec(), vec![2, 2]);
}
