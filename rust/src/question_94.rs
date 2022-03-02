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
    pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        Self::helper(&root)
    }

    fn helper(root: &Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        if root.is_none() {
            return Vec::new();
        }

        let node = root.as_ref().unwrap().borrow();
        let left = &node.left;
        let val = node.val;
        let right = &node.right;
        vec![Self::helper(left), vec![val], Self::helper(right)].concat()
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::inorder_traversal(TreeNode::new_by_vec(vec![Some(1), None, Some(2), Some(3)])),
        vec![1, 3, 2]
    );
}
