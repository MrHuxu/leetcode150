struct Solution;

impl Solution {
    pub fn generate_parenthesis(n: i32) -> Vec<String> {
        let mut ret: Vec<String> = Vec::new();
        for arr in Self::helper(0, 0, n) {
            ret.push(String::from_utf8_lossy(&arr).to_string());
        }
        ret
    }

    fn helper(cnt_l: i32, cnt_r: i32, n: i32) -> Vec<Vec<u8>> {
        if cnt_l == n {
            return vec![vec![')' as u8; (n - cnt_r) as usize]];
        }

        let mut ret: Vec<Vec<u8>> = Vec::new();

        if cnt_l < n {
            for next in Self::helper(cnt_l + 1, cnt_r, n) {
                ret.push(vec![vec!['(' as u8], next].concat());
            }
        }
        if cnt_r < cnt_l {
            for next in Self::helper(cnt_l, cnt_r + 1, n) {
                ret.push(vec![vec![')' as u8], next].concat());
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(Solution::generate_parenthesis(1), vec![String::from("()")]);
    assert_eq!(
        Solution::generate_parenthesis(3),
        vec![
            String::from("((()))"),
            String::from("(()())"),
            String::from("(())()"),
            String::from("()(())"),
            String::from("()()()")
        ]
    );
}
