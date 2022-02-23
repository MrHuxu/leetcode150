struct Solution;

impl Solution {
    pub fn convert(s: String, num_rows: i32) -> String {
        if num_rows == 1 {
            return s;
        }

        let mut rows: Vec<Vec<u8>> = vec![vec![]; num_rows as usize];

        for (i, ch) in s.as_bytes().iter().enumerate() {
            let tmp = i % (num_rows as usize * 2 - 2);
            if tmp < num_rows as usize {
                rows[tmp].push(*ch);
            } else {
                rows[num_rows as usize - (tmp - num_rows as usize + 1) - 1].push(*ch);
            }
        }

        let mut bytes: Vec<u8> = Vec::new();
        for row in rows.iter() {
            for ch in row.iter() {
                bytes.push(*ch);
            }
        }

        String::from_utf8_lossy(&bytes).to_string()
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::convert(String::from("PAYPALISHIRING"), 3),
        String::from("PAHNAPLSIIGYIR")
    );
    assert_eq!(
        Solution::convert(String::from("PAYPALISHIRING"), 4),
        String::from("PINALSIGYAHRPI")
    );
    assert_eq!(Solution::convert(String::from("A"), 1), String::from("A"));
}
