struct Solution;

impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut ret = 0;

        for i in 1..prices.len() {
            if prices[i] > prices[i - 1]  {
                ret += prices[i] - prices[i - 1];
            }
        }
        
        ret
    }
}

#[test]
fn test() {
    assert_eq!(Solution::max_profit(vec![7,1,5,3,6,4]), 7);
}
