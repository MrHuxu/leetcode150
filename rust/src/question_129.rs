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
    pub fn sum_numbers(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        Self::helper(root.as_ref(), 0)
    }

    fn helper(root: Option<&Rc<RefCell<TreeNode>>>, pre_sum: i32) -> i32 {
        let borrowed_node = root.unwrap().borrow();
        
        let sum = pre_sum * 10 + borrowed_node.val;

        if borrowed_node.left.is_none() && borrowed_node.right.is_none() {
            return sum;
        }

        let left_sum = if borrowed_node.left.is_some() {
            Self::helper(borrowed_node.left.as_ref(), sum)
        } else {
            0
        };
        let right_sum = if borrowed_node.right.is_some() {
            Self::helper(borrowed_node.right.as_ref(), sum)
        } else {
            0
        };

        left_sum + right_sum
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::sum_numbers(TreeNode::new_by_vec(vec![Some(1), Some(2), Some(3)])),
        25
    );
}
