struct Solution;

impl Solution {
    pub fn full_justify(words: Vec<String>, max_width: i32) -> Vec<String> {
        let mut ret = Vec::new();

        let mut i = 0;
        while i < words.len() {
            let mut total_len = words[i].len();
            let mut words_cnt = 1;
            let mut total_words_len = words[i].len();
            for j in i + 1..words.len() {
                if total_len + 1 + words[j].len() <= max_width as usize {
                    total_len += 1 + words[j].len();
                    total_words_len += words[j].len();
                    words_cnt += 1;
                } else {
                    break;
                }
            }

            let mut line = words[i].clone();
            if words_cnt == 1 {
                for _j in 0..(max_width as usize - total_len) {
                    line.push_str(" ")
                }
            } else if i + words_cnt == words.len() {
                for j in 1..words_cnt {
                    line.push_str(" ");
                    line.push_str(words[i + j].as_str());
                }
                for _j in 0..(max_width as usize - total_words_len - (words_cnt - 1)) {
                    line.push_str(" ")
                }
            } else {
                let space_len = (max_width as usize - total_words_len) / (words_cnt - 1);
                let large_space_count =
                    max_width as usize - total_words_len - (space_len * (words_cnt - 1));

                for j in 1..words_cnt {
                    for _k in 0..(space_len + if j <= large_space_count { 1 } else { 0 }) {
                        line.push_str(" ")
                    }
                    line.push_str(words[i + j].as_str());
                }
            }

            ret.push(line);
            i += words_cnt;
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::full_justify(
            vec![
                String::from("This"),
                String::from("is"),
                String::from("an"),
                String::from("example"),
                String::from("of"),
                String::from("text"),
                String::from("justification.")
            ],
            16
        ),
        vec![
            String::from("This    is    an"),
            String::from("example  of text"),
            String::from("justification.  ")
        ]
    );
    assert_eq!(
        Solution::full_justify(
            vec![
                String::from("What"),
                String::from("must"),
                String::from("be"),
                String::from("acknowledgment"),
                String::from("shall"),
                String::from("be")
            ],
            16
        ),
        vec![
            String::from("What   must   be"),
            String::from("acknowledgment  "),
            String::from("shall be        ")
        ]
    );
    assert_eq!(
        Solution::full_justify(
            vec![
                String::from("Science"),
                String::from("is"),
                String::from("what"),
                String::from("we"),
                String::from("understand"),
                String::from("well"),
                String::from("enough"),
                String::from("to"),
                String::from("explain"),
                String::from("to"),
                String::from("a"),
                String::from("computer."),
                String::from("Art"),
                String::from("is"),
                String::from("everything"),
                String::from("else"),
                String::from("we"),
                String::from("do")
            ],
            20
        ),
        vec![
            String::from("Science  is  what we"),
            String::from("understand      well"),
            String::from("enough to explain to"),
            String::from("a  computer.  Art is"),
            String::from("everything  else  we"),
            String::from("do                  ")
        ]
    );

    assert_eq!(
        Solution::full_justify(
            vec![
                String::from("ask"),
                String::from("not"),
                String::from("what"),
                String::from("your"),
                String::from("country"),
                String::from("can"),
                String::from("do"),
                String::from("for"),
                String::from("you"),
                String::from("ask"),
                String::from("what"),
                String::from("you"),
                String::from("can"),
                String::from("do"),
                String::from("for"),
                String::from("your"),
                String::from("country")
            ],
            16
        ),
        vec![
            String::from("ask   not   what"),
            String::from("your country can"),
            String::from("do  for  you ask"),
            String::from("what  you can do"),
            String::from("for your country")
        ]
    );
}
