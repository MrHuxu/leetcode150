struct Solution;

use std::cmp::max;
use std::i32::MIN;
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        if prices.len() < 2 {
            return 0;
        }

        let (mut buy1, mut sell1, mut buy2, mut sell2) = (MIN, 0, MIN, 0);
        for price in prices {
            buy1 = max(buy1, -price);
            sell1 = max(sell1, buy1 + price);
            buy2 = max(buy2, sell1 - price);
            sell2 = max(sell2, buy2 + price);
        }

        sell2
    }
}

#[test]
fn name() {
    assert_eq!(Solution::max_profit(vec![3, 3, 5, 0, 0, 3, 1, 4]), 6);
    assert_eq!(Solution::max_profit(vec![1, 2, 3, 4, 5]), 4);
}

