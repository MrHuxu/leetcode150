pub struct Solution;

use std::{
    cmp,
    collections::{HashMap, HashSet},
};

impl Solution {
    pub fn longest_consecutive(nums: Vec<i32>) -> i32 {
        let mut ret = 0;

        let mut map_head_to_tail = HashMap::new();
        let mut map_tail_to_head = HashMap::new();
        let mut visited = HashSet::new();
        for num in nums {
            if visited.contains(&num) {
                continue;
            } else {
                visited.insert(num);
            }

            match (
                map_head_to_tail.get(&(num + 1)).copied(),
                map_tail_to_head.get(&(num - 1)).copied(),
            ) {
                (Some(tail), Some(head)) => {
                    ret = cmp::max(ret, tail - head + 1);
                    map_head_to_tail.insert(head, tail);
                    map_tail_to_head.insert(tail, head);
                }
                (Some(tail), _) => {
                    ret = cmp::max(ret, tail - num + 1);
                    map_head_to_tail.insert(num, tail);
                    map_tail_to_head.insert(tail, num);
                }
                (_, Some(head)) => {
                    ret = cmp::max(ret, num - head + 1);
                    map_head_to_tail.insert(head, num);
                    map_tail_to_head.insert(num, head);
                }
                (_, _) => {
                    ret = cmp::max(ret, 1);
                    map_head_to_tail.insert(num, num);
                    map_tail_to_head.insert(num, num);
                }
            }
        }

        ret
    }
}

#[test]
fn test() {
    assert_eq!(4, Solution::longest_consecutive(vec![100, 4, 200, 1, 3, 2]));
    assert_eq!(
        9,
        Solution::longest_consecutive(vec![0, 3, 7, 2, 5, 8, 4, 6, 0, 1])
    );
}
