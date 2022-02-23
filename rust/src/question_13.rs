struct Solution;

impl Solution {
    pub fn roman_to_int(s: String) -> i32 {
        let mut ret = 0;
        let bytes = s.as_bytes();
        for i in 0..bytes.len() {
            ret += match bytes[i] as char {
                'I' => 1,
                'V' => {
                    if i > 0 && bytes[i - 1] == 'I' as u8 {
                        3
                    } else {
                        5
                    }
                }
                'X' => {
                    if i > 0 && bytes[i - 1] == 'I' as u8 {
                        8
                    } else {
                        10
                    }
                }
                'L' => {
                    if i > 0 && bytes[i - 1] == 'X' as u8 {
                        30
                    } else {
                        50
                    }
                }
                'C' => {
                    if i > 0 && bytes[i - 1] == 'X' as u8 {
                        80
                    } else {
                        100
                    }
                }
                'D' => {
                    if i > 0 && bytes[i - 1] == 'C' as u8 {
                        300
                    } else {
                        500
                    }
                }
                _ => {
                    if i > 0 && bytes[i - 1] == 'C' as u8 {
                        800
                    } else {
                        1000
                    }
                }
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(Solution::roman_to_int(String::from("III")), 3);
    assert_eq!(Solution::roman_to_int(String::from("LVIII")), 58);
    assert_eq!(Solution::roman_to_int(String::from("MCMXCIV")), 1994);
}
