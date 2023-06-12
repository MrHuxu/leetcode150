struct Solution;

impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        if strs.is_empty() {
            return String::new();
        }

        let mut ret = strs[0].clone();
        for i in 1..strs.len() {
            while !strs[i].starts_with(&ret) {
                ret = ret[..ret.len() - 1].to_string();
            }
        }
        ret
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::longest_common_prefix(vec![String::from("cir"), String::from("car")]),
        String::from("c")
    );
    assert_eq!(
        Solution::longest_common_prefix(vec![
            String::from("flower"),
            String::from("flow"),
            String::from("flight")
        ]),
        String::from("fl")
    );
    assert_eq!(
        Solution::longest_common_prefix(vec![
            String::from("dog"),
            String::from("racecar"),
            String::from("car")
        ]),
        String::from("")
    );
}
