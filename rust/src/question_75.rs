struct Solution;

impl Solution {
    pub fn sort_colors(nums: &mut Vec<i32>) {
        let mut idx_0 = 0;
        let mut idx_1 = 0;
        let mut idx_2 = 0;
        for i in 0..nums.len() {
            match nums[i] {
                0 => {
                    nums[idx_2] = 2;
                    idx_2 += 1;
                    nums[idx_1] = 1;
                    idx_1 += 1;
                    nums[idx_0] = 0;
                    idx_0 += 1;
                }
                1 => {
                    nums[idx_2] = 2;
                    idx_2 += 1;
                    nums[idx_1] = 1;
                    idx_1 += 1;
                }
                _ => {
                    nums[idx_2] = 2;
                    idx_2 += 1;
                }
            }
        }
    }
}

#[test]
fn test() {
    let mut nums = vec![2, 0, 2, 1, 1, 0];
    Solution::sort_colors(&mut nums);
    assert_eq!(nums, vec![0, 0, 1, 1, 2, 2]);

    nums = vec![2, 0, 1];
    Solution::sort_colors(&mut nums);
    assert_eq!(nums, vec![0, 1, 2]);
}
