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
    pub fn is_symmetric(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if root.is_none() {
            return true;
        }

        let root_node = root.as_ref().unwrap().borrow();
        Self::helper(root_node.left.as_ref(), root_node.right.as_ref())
    }

    fn helper(p: Option<&Rc<RefCell<TreeNode>>>, q: Option<&Rc<RefCell<TreeNode>>>) -> bool {
        match (p, q) {
            (None, None) => true,
            (None, Some(_val)) | (Some(_val), None) => false,
            (Some(p_val), Some(q_val)) => {
                let p_node = p_val.borrow();
                let q_node = q_val.borrow();

                p_node.val == q_node.val
                    && Self::helper(p_node.left.as_ref(), q_node.right.as_ref())
                    && Self::helper(p_node.right.as_ref(), q_node.left.as_ref())
            }
        }
    }
}

#[test]
fn test() {
    assert_eq!(
        Solution::is_symmetric(TreeNode::new_by_vec(vec![
            Some(1),
            Some(2),
            Some(2),
            Some(3),
            Some(4),
            Some(4),
            Some(3)
        ])),
        true
    );
    assert_eq!(
        Solution::is_symmetric(TreeNode::new_by_vec(vec![
            Some(1),
            Some(2),
            Some(2),
            None,
            Some(3),
            None,
            Some(3)
        ])),
        false
    );
}
