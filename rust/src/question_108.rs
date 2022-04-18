use super::tree::*;

struct Solution;

// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn sorted_array_to_bst(nums: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        if nums.is_empty() {
            return None;
        }

        let idx = nums.len() / 2;
        Some(Rc::new(RefCell::new(TreeNode {
            val: nums[idx],
            left: Self::sorted_array_to_bst(nums[..idx].to_vec()),
            right: Self::sorted_array_to_bst(nums[idx + 1..].to_vec()),
        })))
    }
}

#[test]
fn name() {
    assert_eq!(
        Solution::sorted_array_to_bst(vec![-10, -3, 0, 5, 9]),
        TreeNode::new_by_vec(vec![
            Some(0),
            Some(-3),
            Some(9),
            Some(-10),
            None,
            Some(5)
        ])
    );
}
