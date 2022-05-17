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
    pub fn has_path_sum(root: Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> bool {
        Self::helper(&root, target_sum)
    }

    pub fn helper(root: &Option<Rc<RefCell<TreeNode>>>, target_sum: i32) -> bool {
        if let Some(node) = root.as_ref() {
            let borrowed_node = node.borrow();

            if borrowed_node.left.is_none()
                && borrowed_node.right.is_none()
                && borrowed_node.val == target_sum
            {
                return true;
            }

            Self::helper(&borrowed_node.left, target_sum - borrowed_node.val)
                || Self::helper(&borrowed_node.right, target_sum - borrowed_node.val)
        } else {
            false
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::has_path_sum(
            TreeNode::new_by_vec(vec![
                Some(5),
                Some(4),
                Some(8),
                Some(11),
                None,
                Some(13),
                Some(4),
                Some(7),
                Some(2),
                None,
                None,
                None,
                Some(1)
            ]),
            22
        ),
        true
    );
    assert_eq!(
        Solution::has_path_sum(TreeNode::new_by_vec(vec![Some(1), Some(2), Some(3),]), 5),
        false
    );
    assert_eq!(Solution::has_path_sum(None, 0), false);
}
