struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn int_to_roman(num: i32) -> String {
        let mut sym_map: HashMap<i32, u8> = HashMap::new();
        sym_map.insert(1, 'I' as u8);
        sym_map.insert(5, 'V' as u8);
        sym_map.insert(10, 'X' as u8);
        sym_map.insert(50, 'L' as u8);
        sym_map.insert(100, 'C' as u8);
        sym_map.insert(500, 'D' as u8);
        sym_map.insert(1000, 'M' as u8);

        let mut ret: Vec<u8> = Vec::new();
        let mut mut_num = num;
        for num in vec![1000, 100, 10, 1] {
            if mut_num < num {
                continue;
            }

            let times = mut_num / num;
            mut_num -= num * times;

            let sym = sym_map.get(&num).unwrap();
            let sym_5 = sym_map.get(&(num * 5)).unwrap_or(&0);
            let sym_10 = sym_map.get(&(num * 10)).unwrap_or(&0);

            match times {
                0 | 1 | 2 | 3 => {
                    for _i in 0..times {
                        ret.push(*sym);
                    }
                }
                4 => {
                    ret.push(*sym);
                    ret.push(*sym_5);
                }
                5 | 6 | 7 | 8 => {
                    ret.push(*sym_5);
                    for _i in 0..times - 5 {
                        ret.push(*sym);
                    }
                }
                _ => {
                    ret.push(*sym);
                    ret.push(*sym_10);
                }
            }
        }

        String::from_utf8_lossy(&ret).to_string()
    }
}

#[test]
fn test() {
    assert_eq!(Solution::int_to_roman(3), String::from("III"));
    assert_eq!(Solution::int_to_roman(58), String::from("LVIII"));
    assert_eq!(Solution::int_to_roman(1994), String::from("MCMXCIV"));
}
