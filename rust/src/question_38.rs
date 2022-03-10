struct Solution;

impl Solution {
    pub fn count_and_say(n: i32) -> String {
        let mut ret = vec!['1' as u8];

        for _ in 2..=n {
            let mut next_ret = Vec::new();

            let mut ch = 48;
            let mut cnt = 0;
            for num in ret.iter() {
                if *num == ch {
                    cnt += 1;
                    continue;
                }

                if ch != 48 {
                    next_ret.push(cnt + 48);
                    next_ret.push(ch);
                }

                ch = *num;
                cnt = 1;
            }
            next_ret.push(cnt + 48);
            next_ret.push(ch);

            ret = next_ret;
        }

        String::from_utf8_lossy(&ret).to_string()
    }
}

#[test]
fn test() {
    assert_eq!(Solution::count_and_say(1), String::from("1"));
    assert_eq!(Solution::count_and_say(4), String::from("1211"));
}
