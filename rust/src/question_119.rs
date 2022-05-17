struct Solution;

impl Solution {
    pub fn get_row(row_index: i32) -> Vec<i32> {
        let mut row = vec![1];

        for i in 1..=row_index as usize {
            let mut next_row = vec![1; i + 1];
            for j in 1..i {
                next_row[j] = row[j - 1] + row[j];
            }
            row = next_row;
        }

        row
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::get_row(3),
            vec![1, 3, 3, 1],
    )
}
