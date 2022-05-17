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
use std::cmp::min;
use std::rc::Rc;
impl Solution {
    pub fn min_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if root.is_none() {
            return 0;
        }

        Self::helper(&root).unwrap()
    }

    fn helper(root: &Option<Rc<RefCell<TreeNode>>>) -> Option<i32> {
        if let Some(node) = root.as_ref() {
            Some(1 + match (Self::helper(&node.borrow().left), Self::helper(&node.borrow().right)) {
                (None, None) => 0,
                (Some(val), None) => val,
                (None, Some(val)) => val,
                (Some(left_val), Some(right_val)) => min(left_val, right_val)
            })
        } else {
            None
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::min_depth(TreeNode::new_by_vec(vec![
            Some(3),
            Some(9),
            Some(20),
            None,
            None,
            Some(15),
            Some(7)
        ])),
        2
    );
    assert_eq!(
        Solution::min_depth(TreeNode::new_by_vec(vec![
            Some(1),
            None,
            Some(3),
            None,
            Some(4),
            None,
            Some(5),
            None,
            Some(6),
        ])),
        5
    );
}
