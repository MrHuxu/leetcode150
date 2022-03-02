struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn find_substring(s: String, words: Vec<String>) -> Vec<i32> {
        let word_len = words[0].len();
        let words_cnt = words.len();
        if word_len * words_cnt > s.len() {
            return Vec::new();
        }
        let mut visited = HashMap::new();
        for word in words {
            visited.entry(word).and_modify(|cnt| *cnt += 1).or_insert(1);
        }

        let mut ret = Vec::new();
        for i in 0..=s.len() - word_len * words_cnt {
            if Self::dfs(i, words_cnt, &s, word_len, &mut visited) {
                ret.push(i as i32);
            }
        }
        ret
    }

    fn dfs(
        idx: usize,
        left: usize,
        s: &String,
        word_len: usize,
        visited: &mut HashMap<String, i32>,
    ) -> bool {
        if left == 0 {
            return true;
        }

        let word = s[idx..idx + word_len].to_string();
        if !visited.contains_key(&word) || visited.get(&word).unwrap() <= &0 {
            return false;
        }

        visited.entry(word.clone()).and_modify(|cnt| *cnt -= 1);
        let ret = Self::dfs(idx + word_len, left - 1, s, word_len, visited);
        visited.entry(word).and_modify(|cnt| *cnt += 1);

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::find_substring(
            String::from("a"),
            vec![String::from("a"), String::from("a")]
        ),
        vec![]
    );
    assert_eq!(
        Solution::find_substring(
            String::from("barfoothefoobarman"),
            vec![String::from("foo"), String::from("bar")]
        ),
        vec![0, 9]
    );
    assert_eq!(
        Solution::find_substring(
            String::from("barfoofoobarthefoobarman"),
            vec![
                String::from("bar"),
                String::from("foo"),
                String::from("the")
            ]
        ),
        vec![6, 9, 12]
    );
    assert_eq!(
        Solution::find_substring(
            String::from("barfoofoobarthefoobarman"),
            vec![
                String::from("bar"),
                String::from("foo"),
                String::from("the")
            ]
        ),
        vec![6, 9, 12]
    );
    assert_eq!(
        Solution::find_substring(
            String::from("wordgoodgoodgoodbestword"),
            vec![
                String::from("word"),
                String::from("good"),
                String::from("best"),
                String::from("good")
            ]
        ),
        vec![8]
    )
}
