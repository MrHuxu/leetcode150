struct Solution;

use std::cmp::max;
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut ret = 0;

        let mut min_price = prices[0];
        for price in prices {
            if price > min_price {
                ret = max(ret, price - min_price);
            } else if price < min_price {
                min_price = price;
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(Solution::max_profit(vec![7, 1, 5, 3, 6, 4]), 5);
}
