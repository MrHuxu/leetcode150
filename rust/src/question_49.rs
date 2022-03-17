struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut char_counts = Vec::new();
        for s in strs.iter() {
            let char_count = Self::get_char_count(&s.as_bytes());
            char_counts.push(char_count);
        }

        let mut used_char_counts = Vec::new();
        let mut ret: Vec<Vec<String>> = Vec::new();
        for i in 0..strs.len() {
            let mut matched = false;
            for (j, anagram) in used_char_counts.iter().enumerate() {
                if Self::compare_char_count(&char_counts[i], *anagram) {
                    ret[j].push(strs[i].clone());
                    matched = true;
                    break;
                }
            }

            if !matched {
                ret.push(vec![strs[i].clone()]);
                used_char_counts.push(&char_counts[i]);
            }
        }

        ret
    }

    fn get_char_count(bytes: &[u8]) -> HashMap<u8, i32> {
        let mut count = HashMap::new();

        for byte in bytes {
            let count = count.entry(*byte).or_insert(0);
            *count += 1;
        }

        count
    }

    fn compare_char_count(char_count1: &HashMap<u8, i32>, char_count2: &HashMap<u8, i32>) -> bool {
        let keys1 = char_count1.keys();
        let keys2 = char_count2.keys();
        if keys1.len() != keys2.len() {
            return false;
        }

        for key in keys1 {
            if char_count1.get(key).unwrap() != char_count2.get(key).unwrap_or(&0) {
                return false;
            }
        }

        true
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::group_anagrams(vec![
            String::from("eat"),
            String::from("tea"),
            String::from("tan"),
            String::from("ate"),
            String::from("nat"),
            String::from("bat")
        ]),
        vec![
            vec![
                String::from("eat"),
                String::from("tea"),
                String::from("ate")
            ],
            vec![String::from("tan"), String::from("nat")],
            vec![String::from("bat")]
        ]
    );
}
