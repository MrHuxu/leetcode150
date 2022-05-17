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
use std::i32::MIN;
use std::rc::Rc;
impl Solution {
    pub fn max_path_sum(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut ret = MIN;
        Self::helper(root.as_ref(), &mut ret);
        ret
    }

    fn helper(root: Option<&Rc<RefCell<TreeNode>>>, ret: &mut i32) -> i32 {
        if let Some(node) = root {
            let borrowed_node = node.borrow();
            let left_sum = Self::helper(borrowed_node.left.as_ref(), ret);
            let right_sum = Self::helper(borrowed_node.right.as_ref(), ret);

            *ret = max(*ret, borrowed_node.val);
            *ret = max(*ret, borrowed_node.val + left_sum + right_sum);

            max(0, max(left_sum, right_sum) + borrowed_node.val)
        } else {
            0
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::max_path_sum(TreeNode::new_by_vec(vec![Some(1), Some(2), Some(3)])),
        6
    );
    assert_eq!(
        Solution::max_path_sum(TreeNode::new_by_vec(vec![Some(-3)])),
        -3
    );
}
