struct Solution;

impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        let mut ret = strs[0].as_bytes().to_vec();

        for str in strs[1..].iter() {
            let bytes = str.as_bytes();

            let mut idx = -1;
            for i in 0..bytes.len() {
                if i < ret.len() && i < bytes.len() && ret[i] == bytes[i] {
                    idx = i as i32;
                } else {
                    break;
                }
            }

            ret = ret[0..(idx + 1) as usize].to_vec();
        }

        String::from_utf8_lossy(&ret).to_string()
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
