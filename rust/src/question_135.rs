pub struct Solution;

impl Solution {
    pub fn candy(ratings: Vec<i32>) -> i32 {
        let mut ret = 0;

        let mut candies = vec![0; ratings.len()];
        for i in 0..candies.len() {
            if i > 0 && ratings[i] > ratings[i - 1] {
                candies[i] = candies[i - 1] + 1;
                ret += candies[i];
                continue;
            }

            candies[i] = 1;
            ret += 1;
        }
        for i in (0..candies.len() - 1).rev() {
            if ratings[i] > ratings[i + 1] && candies[i] <= candies[i + 1] {
                ret += candies[i + 1] + 1 - candies[i];
                candies[i] = candies[i + 1] + 1;
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(5, Solution::candy(vec![1, 0, 2]));
    assert_eq!(4, Solution::candy(vec![1, 2, 2]));
    assert_eq!(13, Solution::candy(vec![1, 2, 87, 87, 87, 2, 1]));
}
