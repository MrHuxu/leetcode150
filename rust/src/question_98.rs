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
    pub fn is_valid_bst(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        Self::helper(root.as_ref(), None, None)
    }

    fn helper(
        root: Option<&Rc<RefCell<TreeNode>>>,
        min_val: Option<i32>,
        max_val: Option<i32>,
    ) -> bool {
        if root.is_none() {
            return true;
        }

        let node = root.unwrap().borrow();

        if min_val.is_some() && node.val < min_val.unwrap() {
            return false;
        }
        if max_val.is_some() && node.val > max_val.unwrap() {
            return false;
        }

        Self::helper(node.left.as_ref(), min_val, Some(node.val))
            && Self::helper(node.right.as_ref(), Some(node.val), max_val)
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::is_valid_bst(TreeNode::new_by_vec(vec![Some(2), Some(1), Some(3)])),
        true
    );
    assert_eq!(
        Solution::is_valid_bst(TreeNode::new_by_vec(vec![
            Some(5),
            Some(1),
            Some(4),
            None,
            None,
            Some(3),
            Some(6)
        ])),
        false
    );
    assert_eq!(
        Solution::is_valid_bst(TreeNode::new_by_vec(vec![
            Some(5),
            Some(4),
            Some(6),
            None,
            None,
            Some(3),
            Some(7)
        ])),
        false
    );
}
