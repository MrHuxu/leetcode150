struct Solution;

impl Solution {
    pub fn multiply(num1: String, num2: String) -> String {
        if num1 == String::from("0") || num2 == String::from("0") {
            return String::from("0");
        }

        let mut nums = Vec::new();

        for (i, ch2) in num2.as_bytes().iter().rev().enumerate() {
            let mut num = Vec::new();

            for _ in 0..i {
                num.push(0);
            }

            let mut carry = 0;
            for ch1 in num1.as_bytes().iter().rev() {
                let tmp = (*ch2 - '0' as u8) as i32 * (*ch1 - '0' as u8) as i32 + carry;
                num.push(tmp % 10);
                carry = tmp / 10;
            }
            if carry != 0 {
                num.push(carry);
            }

            nums.push(num);
        }

        let mut ret = Vec::new();

        let mut idx = 0;
        let mut carry = 0;
        loop {
            let mut should_break = true;

            let mut sum = carry;
            for num in nums.iter() {
                if idx >= num.len() {
                    continue;
                }

                should_break = false;
                sum += num[idx];
            }

            if should_break {
                break;
            }

            ret.push((sum % 10) as u8 + '0' as u8);
            carry = sum / 10;
            idx += 1;
        }
        if carry > 0 {
            ret.push(carry as u8 + '0' as u8);
        }

        ret.reverse();
        String::from_utf8_lossy(&ret).to_string()
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::multiply(String::from("2"), String::from("3")),
        String::from("6")
    );
    assert_eq!(
        Solution::multiply(String::from("123"), String::from("456")),
        String::from("56088")
    );
    assert_eq!(
        Solution::multiply(String::from("9"), String::from("9")),
        String::from("81")
    );
    assert_eq!(
        Solution::multiply(String::from("123456789"), String::from("987654321")),
        String::from("121932631112635269")
    );
}
