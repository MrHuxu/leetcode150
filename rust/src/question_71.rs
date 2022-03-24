struct Solution;

impl Solution {
    pub fn simplify_path(path: String) -> String {
        let strs = path.split("/");

        let mut stack = Vec::new();
        for s in strs {
            match s {
                "" | "." => continue,
                ".." => {
                    if stack.len() > 0 {
                        stack.pop();
                    }
                }
                _ => stack.push(String::from(s)),
            };
        }

        String::from("/") + &stack.join("/")
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::simplify_path(String::from("/home/")),
        String::from("/home")
    );
    assert_eq!(
        Solution::simplify_path(String::from("/../")),
        String::from("/")
    );
    assert_eq!(
        Solution::simplify_path(String::from("/home//foo/")),
        String::from("/home/foo")
    );
}
