struct Solution;

impl Solution {
    pub fn combine(n: i32, k: i32) -> Vec<Vec<i32>> {
        Self::dfs(k, 1, n)
    }

    fn dfs(remaining: i32, curr: i32, n: i32) -> Vec<Vec<i32>> {
        if remaining == 0 {
            return vec![vec![]];
        }

        let mut ret = Vec::new();

        for num in curr..=n {
            for tmp in Self::dfs(remaining - 1, num + 1, n) {
                ret.push(vec![vec![num], tmp].concat());
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::combine(4, 2),
        vec![
            vec![1, 2],
            vec![1, 3],
            vec![1, 4],
            vec![2, 3],
            vec![2, 4],
            vec![3, 4]
        ]
    );
    assert_eq!(Solution::combine(1, 1), vec![vec![1]]);
}
