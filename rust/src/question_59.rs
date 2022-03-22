struct Solution;

impl Solution {
    pub fn generate_matrix(n: i32) -> Vec<Vec<i32>> {
        let mut ret = vec![vec![0; n as usize]; n as usize];
        let mut num = 1;
        let n = n as usize;
        for i in 0..(n + 1) / 2 {
            for j in i..n - i {
                ret[i][j] = num;
                num += 1;
            }

            for j in i + 1..n - 1 - i {
                ret[j][n - 1 - i] = num;
                num += 1;
            }

            if n - 1 - i == i {
                continue;
            }

            for j in i..n - i {
                ret[n - 1 - i][n - 1 - j] = num;
                num += 1;
            }

            for j in i + 1..n - 1 - i {
                ret[n - 1 - j][i] = num;
                num += 1;
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::generate_matrix(3),
        vec![vec![1, 2, 3], vec![8, 9, 4], vec![7, 6, 5]]
    )
}
