struct Solution;

impl Solution {
    pub fn str_str(haystack: String, needle: String) -> i32 {
        if haystack.len() < needle.len() {
            return -1;
        }
        if haystack == needle {
            return 0;
        }

        let bytes_h = haystack.as_bytes();
        let bytes_n = needle.as_bytes();
        for i in 0..=bytes_h.len() - bytes_n.len() {
            let mut is_match = true;

            for j in 0..bytes_n.len() {
                if bytes_h[i + j] != bytes_n[j] {
                    is_match = false;
                    break;
                }
            }

            if is_match {
                return i as i32;
            }
        }

        -1
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::str_str(String::from("hello"), String::from("ll")),
        2
    );
    assert_eq!(
        Solution::str_str(String::from("aaaaa"), String::from("bba")),
        -1
    );
    assert_eq!(Solution::str_str(String::from(""), String::from("")), 0);
}
