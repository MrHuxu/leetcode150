struct Solution;

impl Solution {
    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        Self::dfs(vec![], 0, &nums)
    }

    fn dfs(pre: Vec<i32>, idx: usize, nums: &Vec<i32>) -> Vec<Vec<i32>> {
        let mut ret = vec![pre.clone()];

        for i in idx..nums.len() {
            let mut next = pre.clone();
            next.push(nums[i]);
            for tmp in Self::dfs(next, i + 1, nums) {
                ret.push(tmp);
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::subsets(vec![1, 2, 3]),
        vec![
            vec![],
            vec![1],
            vec![1, 2],
            vec![1, 2, 3],
            vec![1, 3],
            vec![2],
            vec![2, 3],
            vec![3]
        ]
    );
    assert_eq!(Solution::subsets(vec![0]), vec![vec![], vec![0]]);
}
