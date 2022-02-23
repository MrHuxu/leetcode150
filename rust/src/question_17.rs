struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn letter_combinations(digits: String) -> Vec<String> {
        let mut board: HashMap<u8, Vec<u8>> = HashMap::new();
        board.insert('2' as u8, "abc".as_bytes().to_vec());
        board.insert('3' as u8, "def".as_bytes().to_vec());
        board.insert('4' as u8, "ghi".as_bytes().to_vec());
        board.insert('5' as u8, "jkl".as_bytes().to_vec());
        board.insert('6' as u8, "mno".as_bytes().to_vec());
        board.insert('7' as u8, "pqrs".as_bytes().to_vec());
        board.insert('8' as u8, "tuv".as_bytes().to_vec());
        board.insert('9' as u8, "wxyz".as_bytes().to_vec());

        let mut data: Vec<Vec<u8>> = Vec::new();
        for digit in digits.as_bytes().iter() {
            let bytes = board.get(digit).unwrap();

            if data.len() == 0 {
                for byte in bytes {
                    data.push(vec![*byte]);
                }
                continue;
            }

            let mut next_data: Vec<Vec<u8>> = Vec::new();
            for pre in data {
                for byte in bytes {
                    let mut tmp = pre.clone();
                    tmp.push(*byte);
                    next_data.push(tmp);
                }
            }
            data = next_data;
        }

        let mut ret: Vec<String> = Vec::new();
        for d in data.iter() {
            ret.push(String::from_utf8_lossy(d).to_string());
        }
        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::letter_combinations(String::from("23")),
        vec![
            String::from("ad"),
            String::from("ae"),
            String::from("af"),
            String::from("bd"),
            String::from("be"),
            String::from("bf"),
            String::from("cd"),
            String::from("ce"),
            String::from("cf")
        ]
    );
    assert_eq!(
        Solution::letter_combinations(String::from("")),
        Vec::<String>::new()
    );
    assert_eq!(
        Solution::letter_combinations(String::from("2")),
        vec![String::from("a"), String::from("b"), String::from("c")]
    );
}
