struct Solution;

impl Solution {
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        let mut ret = vec![vec![]; num_rows as usize];

        ret[0] = vec![1];
        for i in 1..num_rows as usize {
            ret[i] = vec![1; i + 1];
            for j in 1..i {
                ret[i][j] = ret[i - 1][j - 1] + ret[i - 1][j];
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::generate(5),
        vec![
            vec![1],
            vec![1, 1],
            vec![1, 2, 1],
            vec![1, 3, 3, 1],
            vec![1, 4, 6, 4, 1],
        ]
    )
}
