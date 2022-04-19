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
use std::cmp::max;
use std::rc::Rc;
impl Solution {
    pub fn is_balanced(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        let (_depth, is_balanced) = Self::helper(&root);
        is_balanced
    }

    fn helper(root: &Option<Rc<RefCell<TreeNode>>>) -> (i32, bool) {
        if let Some(node) = root.as_ref() {
            let (left_depth, left_is_balanced) = Self::helper(&node.borrow().left);
            let (right_depth, right_is_balanced) = Self::helper(&node.borrow().right);
            (
                max(left_depth, right_depth) + 1,
                left_is_balanced && right_is_balanced && (left_depth - right_depth).abs() <= 1,
            )
        } else {
            (0, true)
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::is_balanced(TreeNode::new_by_vec(vec![
            Some(3),
            Some(9),
            Some(20),
            None,
            None,
            Some(15),
            Some(7)
        ])),
        true
    );
    assert_eq!(
        Solution::is_balanced(TreeNode::new_by_vec(vec![
            Some(1),
            Some(2),
            Some(2),
            Some(3),
            Some(3),
            None,
            None,
            Some(4),
            Some(4)
        ])),
        false
    );
}
